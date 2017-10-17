## selpg
[selpg](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html) 是一个自定义命令行程序，全称select page，即从源（标准输入流或文件）读取指定页数的内容到目的地（标准输出流或给给打印机打印）

##参数使用说明
详细用法请参考 [C语言实现自定义selpg](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)，我这里是为练习go而写的，与原文用法有差异

selpg [-s startPage] [-e endPage] [-l linePerPage | -f] [-d dest] filename

必需标志以及参数：

 - -s，后面接开始读取的页号 int
 - -e，后面接结束读取的页号 int
 s和e都要大于1，并且s <= e，否则提示错误

可选参数

 - -l，后面跟行数 int，代表多少行分为一页，不指定 -l 又缺少 -f 则默认按照72行分一页
 - -f，该标志无参数，代表按照分页符'\f' 分页
 - -d，~~后面接打印机标号，用于将内容传送给打印机打印~~ 我没有打印机用于测试，所以当使用 -d destination(随便一个字符串作参数)时，就会通过管道把内容发送给 grep命令，并把grep处理结果显示到屏幕
 - filename，唯一一个无标识参数，代表选择读取的文件名

## 设计

 - 结构体记录参数对应的值易于处理
 ```
type selpgArgs struct {
	startPage  int //起始页号
	endPage    int //终止页号
	pageLen    int //每页的行数
	pageType   string//分页类型，默认 -l 按行数分页
	printDest  string//选择的打印机
	inFilename string//读取的文件名
}
 ```

 - 解析命令的标志和对应参数
 
 当我们输入命令 `$ ./selpg -s 1 -e 6 -f -d dest input_file`时，有三种参数需要获取
 1.  类似"-s 30"形式，即 标志 + 对应参数值
      要获取 -s 的参数30，只需调用如下函数
    `flag.IntVar(&sa.startPage, "s", -1, "the start page")`
    上面一行代码定义了标志 s 的默认参数值是 -1，当接受到 "-s 30" 则会把 30 赋值给 sa.startPage，否则把 -1 给sa.startPage
    
 2. 类似"-f"形式，单独一个标志，不需对应参数值
     我们也能通过flag.TypeVar(*val, flag, defaultVal, usage)来定义 -f，但是因为命令行语法的原因，**要识别 -flag ，那么它对应参数的值只能是 bool ，不能是 string int之类的**。
     因此，采用
     `input_f := flag.Bool("f", false, "some usage of -f")` 
     通过上面函数，当命令行有输入 -f 时，input_f (-f 对应的参数)即为true，否则为false
     
     
 3. 类似"input_file"，没有标志，只有参数
     flag.Arg(i) 则表示命令中的无标识参数中的第i的参数值
     因此，通过 `sa.inFilename = flag.Arg(0)` 即能获取该参数
     
     
 - 分页方式的处理
    1. 按行分页 -l pageLe  
 ```
if sa.pageType == "l" {
		line := bufio.NewScanner(fin)
		//一行一行读取文件
		for line.Scan() {
		    //如果当前的页号满足范围，则把处于该页中的该行输出
			if currPage >= sa.startPage && currPage <= sa.endPage {
				fout.Write([]byte(line.Text() + "\n"))
				if sa.printDest != "" {
					inpipe.Write([]byte(line.Text() + "\n"))
				}
			}
			//满行则换页
			currLine++
			if currLine % sa.pageLen == 0 {
				currPage++
				currLine = 0
			}
		}
	}
 ```
   2. 按分页符分页 -f
        根据分页符，先把一页内容读进
 ```
rd := bufio.NewReader(fin)
		for {
			//ReadString每次从当前位置读取，直到遇到换页符'\f'，然后下次从'\f'后面开始继续读取，即每次读一页
			page, ferr := rd.ReadString('\f')
			if ferr != nil || ferr == io.EOF {
				//注意把最后一页输出
				if ferr == io.EOF {
					if currPage >= sa.startPage && currPage <= sa.endPage {
						fmt.Fprintf(fout, "%s", page)
						currPage++
					}
				}
				break
			}
			page = strings.Replace(page, "\f", "", -1)

			if currPage >= sa.startPage && currPage <= sa.endPage {
				fmt.Fprintf(fout, "%s", page)
				currPage++
			}
		}
 ```

 - 通过管道把selpg输出给另一个命令
   调用函数 `exec.Command("命令名", "该命令的相关参数", ...)` 即可返回获得该命令结构，再调用 cmd.StdinPipe() 可以获得管道的输入管道，然后把我们的 selpg 输出到该管道即可。
   我是把selpg 输出给了命令 grep，查找出selpg的输出中与文件keyword内容相关的部分，再输出到屏幕
 ```
if sa.printDest != "" {
		cmd := exec.Command("grep", "-nf", "keyword")
		inpipe, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer inpipe.Close()
		cmd.Stdout = fout
		cmd.Start()
	}
 ```



## 正确用法示例

下面的例子中我这里只介绍本程序的用法
用到的相关文件：
①input_file为输入文件
②input_file_f为输入文件，经处理过，每五行以一个分页符 '\f' 结束，用于测试 -f 标志
③output_file为输出文件
④error_file为保存程序出错信息的文件
⑤keyword该文件里面只有i一个字母。
①②④内容如下，②③初始为空文件
![这里写图片描述](http://img.blog.csdn.net/20171017200609513?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![input_file_f](http://img.blog.csdn.net/20171018011737363?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![keyword](http://img.blog.csdn.net/20171017211317166?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

1. 不选择文件，则默认从键盘读入，输出到屏幕
 ```
$ ./selpg -s 1 -e 3
 ```
执行上面命令后，当我们从键盘输入一行，就会输出一行到屏幕，而且不用 -l 指定页的行数，所以(默认72行为一页)直到我们输够 3*72 行后，程序才会终止读取……

2. 不指定-l，读取文件，则默认以72行为一页 
 ```
$ ./selpg -s 1 -e 6 input_file
输出：
this
is
my
first
go-lang
program
and
i feel
really
happy
./selpg: end_page (6) greater than total pages (1), less output than expected
 ```
因为input_file只有6行，所以只能算一页，而我们指定了从第一页读取到第六页，因此会提示错误：指定结束页6比文件总页数1少了

3. 从input_file读取其第2—5页的内容，指定每页的行数为 1
 ```
$ ./selpg -s 2 -e 5 -l 1 input_file
输出：
is
my
first
go-lang
 ```

4. 从input_file读取其第2—5页内容，指定每页的行数为 1，把结果写入文件output_file
 ```
$ ./selpg -s 2 -e 5 -l 1 input_file > output_file
 ```
![output_file](http://img.blog.csdn.net/20171017201532760?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

5. 从input_file_f读取第1—2页内容 ，通过分页符定界，结果输出到屏幕
 ```
$ ./selpg -s 1 -e 2 -f input_file_f
输出：
this is my first go-lang program
in fact, i enjoy it
but i have so much homework to do
if i have more time
i can do it better
however, the reson in nonsense
It is not enough for a year.
I want to make my own progress.
So that I could find a great company or graduate
Do your best!hjm!
 ```
可以看到输出10行，原因是原文件中每五行后面才有一个分页符 '\f' 分割 

5. ~~读取文件输出到打印机打印，~~注意我已经修改了该功能用于测试！！因为我没打印机测试，所以把这功能改为输入到 grep 中，从中找出含有 keyword文件(只含有字母 i)中关键字的内容，并且输出到屏幕上
 ```
$ ./selpg -s 1 -e 7 -l 1 -d none input_file
输出：
this
is
my
first
go-lang
program
and
1:this
2:is
4:first
 ```
可以看到输出中既有读取文件的内容，最后也列出根据关键字 i 搜索出的内容，最后三行带有行号的为 grep 执行结果，说明程序中的管道运用是正确的

7. 定向其输入流为文件
 ```
$ ./selpg -s 1 -e 1 < input_file
输出：
this
is
my
first
go-lang
program
and
i feel
really
happy
 ```
 
8. 把输出导向到文件output_file
 ```
$ ./selpg -s 1 -e 1 input_file > output_file
 ```
![output](http://img.blog.csdn.net/20171017212940746?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

9. 把错误信息导向文件error_file
 ```
$ ./selpg -s 1 -e  2> error_file
 ```
![errorfile](http://img.blog.csdn.net/20171017212738806?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

10. 将命令"ps"的输出的第1页到第3页写至selpg的标准输出（屏幕）；命令"ps"可以替换为任意其他linux行命令，selpg的输出也能成为另一个命令的输入。
 ```
$ ps | ./selpg -s 1 -e 3
输出：
  PID TTY          TIME CMD
10011 pts/2    00:00:00 bash
10648 pts/2    00:00:00 ps
10649 pts/2    00:00:00 selpg
./selpg: end_page (3) greater than total pages (1), less output than expected
 ```
 
11. 将selpg的输出传给 cat 命令作为输入执行，cat结果显示在屏幕
 ```
$ ./selpg -s 1 -e 1 input_file | cat -n
输出：
     1  this
     2  is
     3  my
     4  first
     5  go-lang
     6  program
     7  and
     8  i feel
     9  really

    10  happy
 ```

## 一些错误用法示例
命令输入出错时，会提示相应用法

1. 标志缺少参数
 ```
$ ./selpg -s -l
输出：
invalid value "-l" for flag -s: strconv.ParseInt: parsing "-l": invalid syntax
Usage of ./selpg:
  -d string
  -e int
         (default -1)
  -f
  -l int
         (default 72)
  -s int
         (default -1)
 ```


2. 缺少必须标志 -s -l

 ```
./selpg -s 1 input_file
输出：
USAGE: ./selpg [-s start_page] [-e end_page] [ -l lines_per_page | -f ] [ -d dest ] [ in_filename ]
 ```
 
3. 文件不存在或没权限打开

 ```
$ ./selpg -s 1 -e 6 noFile
输出：
selpg: could not open input file "noFile"
open noFile: no such file or directory
USAGE: ./selpg [-s start_page] [-e end_page] [ -l lines_per_page | -f ] [ -d dest ] [ in_filename ]
 ```
