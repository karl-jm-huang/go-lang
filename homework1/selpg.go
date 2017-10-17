package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type selpgArgs struct {
	startPage  int
	endPage    int
	pageLen    int //lines per page
	pageType   string
	printDest  string
	inFilename string
}

func main() {
	sa := new(selpgArgs)

	/** 定义标志参数
	 *  -s -e -l -f -d
	 */
	flag.IntVar(&sa.startPage, "s", -1, "")
	flag.IntVar(&sa.endPage, "e", -1, "")
	flag.IntVar(&sa.pageLen, "l", 72, "")
	flag.StringVar(&sa.printDest, "d", "", "")

	/**检查 -f是否存在，注意 -f 只支持bool类型
	 * 默认的，提供了-flag，则对应的值为true
	 */
	exist_f := flag.Bool("f", false, "")
	//解析命令行参数到定义的flag
	flag.Parse()

	if *exist_f {
		sa.pageType = "f"
		sa.pageLen = -1
	} else {
		sa.pageType = "l"
	}

	// 非标志参数为文件名
	if flag.NArg() == 1 {
		sa.inFilename = flag.Arg(0)
		//fmt.Printf("now in if flag.NArg() == 1\n")
	} else {
		sa.inFilename = ""
	}

	//检查参数合法性
	checkArgs(*sa, flag.NArg())

	//执行命令
	exeSelpg(*sa)
}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUSAGE: ./selpg [-s start_page] [-e end_page] [ -l lines_per_page | -f ] [ -d dest ] [ in_filename ]\n")
}

func checkArgs(sa selpgArgs, notFlagNum int) {
	s_e_ok := sa.startPage <= sa.endPage && sa.startPage >= 1
	num_ok := notFlagNum == 1 || notFlagNum == 0
	l_f_ok := !(sa.pageType == "f" && sa.pageLen != -1)
	if !s_e_ok || !num_ok || !l_f_ok {
		usage()
		os.Exit(1)
	}
}

func exeSelpg(sa selpgArgs) {
	currPage := 1
	currLine := 0

	fin := os.Stdin
	fout := os.Stdout
	var inpipe io.WriteCloser
	var err error

	//确定输入源，是否改为从文件读入
	if sa.inFilename != "" {
		fin, err = os.Open(sa.inFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "selpg: could not open input file \"%s\"\n", sa.inFilename)
			fmt.Println(err)
			usage()
			os.Exit(1)
		}
		defer fin.Close()
	}

	/**确定输出源, 是否打印到打印机
	 * 由于没有打印机测试，用管道接通 grep 作为测试，结果输出到屏幕
	 * selpg内容通过管道输入给 grep, grep从中搜出带有keyword文件的内容
	 */
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

	// -l 按行分页读取输出
	if sa.pageType == "l" {
		line := bufio.NewScanner(fin)
		for line.Scan() {
			if currPage >= sa.startPage && currPage <= sa.endPage {
				fout.Write([]byte(line.Text() + "\n")) //一行行输出
				if sa.printDest != "" {
					inpipe.Write([]byte(line.Text() + "\n"))
				}
			}
			currLine++
			if currLine%sa.pageLen == 0 {
				currPage++
				currLine = 0
			}
		}
	} else { // -f 按分隔符分页读取输出
		rd := bufio.NewReader(fin)
		for {
			//一页一页的读
			page, ferr := rd.ReadString('\f')
			if ferr != nil || ferr == io.EOF {
				if ferr == io.EOF {
					if currPage >= sa.startPage && currPage <= sa.endPage {
						fmt.Fprintf(fout, "%s", page)
					}
				}
				break
			}
			page = strings.Replace(page, "\f", "", -1)
			currPage++
			if currPage >= sa.startPage && currPage <= sa.endPage {
				fmt.Fprintf(fout, "%s", page)
			}
		}

	}

	if currPage < sa.endPage {
		fmt.Fprintf(os.Stderr, "./selpg: end_page (%d) greater than total pages (%d), less output than expected\n", sa.endPage, currPage)
	}
}
