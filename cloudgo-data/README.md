## 构建数据服务
使用 xorm 代替 database/sql 修改 [课堂项目](https://github.com/pmlpml/golang-learning/web/cloudgo-data)

从编程效率、程序结构、服务性能等角度对比 database/sql 与 xorm 实现的异同！ 

## xorm 是否就是实现了 dao 的自动化
我认为xorm实现了dao的自动化，并且使用反射技术，牺牲性能增加易用性。

xorm使得我们编程效率提高，省去设计实现dao接口的工作，程序结构更简单，但是服务性能经过下面测试有所下降。

## 基本web测试
启动服务器
```
$ ./main
[negroni] listening on :8080
```

用 curl POST 一些数据到网站
```
$ curl -d "username=ooo&departname=1" http://localhost:8080/service/userinfo
{
  "UID": 0,
  "UserName": "ooo",
  "DepartName": "1",
  "CreateAt": "2017-12-01T21:28:18.783865292+08:00"
}

```
查询上传的数据
```
curl http://localhost:8080/service/userinfo?userid=
[
  {
    "UID": 0,
    "UserName": "ooo",
    "DepartName": "1",
    "CreateAt": "2017-12-02T05:28:18+08:00"
  }
]

```

## 使用 ab 测试性能

### 原项目database/sql的测试结果

```
$ ab -c 200 -n 1000 http://localhost:8080/service/userinfo?userid=
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        111 bytes

Concurrency Level:      200
Time taken for tests:   0.242 seconds //测试持续时间
Complete requests:      1000
Failed requests:        366
   (Connect: 0, Receive: 0, Length: 366, Exceptions: 0)
Non-2xx responses:      366
Total transferred:      1679835 bytes
HTML transferred:       1558763 bytes
Requests per second:    4137.39 [#/sec] (mean)
Time per request:       48.340 [ms] (mean) //用户平均请求等待时间
Time per request:       0.242 [ms] (mean, across all concurrent requests) //服务器平均请求等待时间
Transfer rate:          6787.25 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.4      0       5
Processing:     5   44  29.2     37     114
Waiting:        5   44  28.9     37     114
Total:          5   45  30.1     37     118

Percentage of the requests served within a certain time (ms)
  50%     37
  66%     55
  75%     64
  80%     71
  90%     93
  95%    109
  98%    112
  99%    113
 100%    118 (longest request)


```

### 改用xorm后测试结果

```
$ ab -c 200 -n 1000 http://localhost:8080/service/userinfo?userid=
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        230 bytes

Concurrency Level:      200
Time taken for tests:   0.262 seconds //测试持续时间
Complete requests:      1000
Failed requests:        421
   (Connect: 0, Receive: 0, Length: 421, Exceptions: 0)
Non-2xx responses:      421
Total transferred:      1993795 bytes
HTML transferred:       1873163 bytes
Requests per second:    3813.66 [#/sec] (mean) 
Time per request:       52.443 [ms] (mean) //用户平均请求等待时间
Time per request:       0.262 [ms] (mean, across all concurrent requests) //服务器平均请求等待时间
Transfer rate:          7425.45 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.8      0       8
Processing:     4   46  31.3     35     140
Waiting:        3   46  31.2     34     140
Total:          5   47  31.5     36     143

Percentage of the requests served within a certain time (ms)
  50%     36
  66%     53
  75%     71
  80%     80
  90%     97
  95%    107
  98%    126
  99%    141
 100%    143 (longest request)


```

相比之下，从用户和服务器等待时间来看，database/sql 比 xorm 的响应速度快一点
