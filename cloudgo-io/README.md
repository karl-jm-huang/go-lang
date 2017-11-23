## 项目要求
实现一个简单的cloudgo应用，主要包括：
	
	1. 静态文件服务
	2. 支持简单的js访问
	3. 提交表单，并输出一个表格
	4. 对 /unknown 给出开发中的提示，返回码 5xx

## 启动
```
$ go run main.go
[negroni] listening on :8080
```

## 静态文件访问

```
[negroni] 2017-11-23T21:19:52+08:00 | 200 |      100.637µs | localhost:8080 | GET /assets/
[negroni] 2017-11-23T21:20:30+08:00 | 200 |      59.316µs | localhost:8080 | GET /assets/css/
[negroni] 2017-11-23T21:20:55+08:00 | 200 |      149.521µs | localhost:8080 | GET /assets/js/
[negroni] 2017-11-23T21:21:09+08:00 | 200 |      75.278µs | localhost:8080 | GET /assets/images/
```
![这里写图片描述](http://img.blog.csdn.net/20171123212221770?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171123211040006?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171123211102877?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171123211127739?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
## 简单js访问
![这里写图片描述](http://img.blog.csdn.net/20171123211401931?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171123211334525?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

## 模板输出
![这里写图片描述](http://img.blog.csdn.net/20171123212551873?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171123211200047?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

## 提交表单，返回表格输出
![这里写图片描述](http://img.blog.csdn.net/20171123211221333?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171123211234759?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)



## 对/unknown的提示
简单定义一个手柄函数输出语句即可
![这里写图片描述](http://img.blog.csdn.net/20171123211317100?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSDEyNTkwNDAwMzI3/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


