## 构建数据服务
使用 xorm 代替 database/sql 修改 [课堂项目](https://github.com/pmlpml/golang-learning/web/cloudgo-data)

从编程效率、程序结构、服务性能等角度对比 database/sql 与 xorm 实现的异同！ 

## xorm 是否就是实现了 dao 的自动化
我认为xorm实现了dao的自动化，并且使用反射技术，牺牲性能增加易用性。

xorm使得我们编程效率提高，省去设计实现dao接口的工作，程序结构更简单，但是服务性能经过下面测试有所下降。

## 使用 ab 测试性能

### 原项目database/sql的测试结果

```
$ ab -n 1000 -c 200 http://localhost:8080/service/userinfo?userid=
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
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


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        111 bytes

Concurrency Level:      200
Time taken for tests:   3.622 seconds
Complete requests:      1000
Failed requests:        14
   (Connect: 0, Receive: 0, Length: 14, Exceptions: 0)
Non-2xx responses:      14
Total transferred:      292218 bytes
HTML transferred:       168330 bytes
Requests per second:    276.09 [#/sec] (mean)
Time per request:       724.393 [ms] (mean)
Time per request:       3.622 [ms] (mean, across all concurrent requests)
Transfer rate:          78.79 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       3
Processing:     3  330 670.3    154    3034
Waiting:        2  326 671.6    151    3034
Total:          4  331 670.3    155    3034

Percentage of the requests served within a certain time (ms)
  50%    155
  66%    223
  75%    253
  80%    273
  90%    348
  95%   3007
  98%   3021
  99%   3024
 100%   3034 (longest request)
Finished 1000 requests
```

### 改用xorm后测试结果

```
$ ab -n 1000 -c 200 http://localhost:8080/service/userinfo?userid=
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
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


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        116 bytes

Concurrency Level:      200
Time taken for tests:   3.873 seconds
Complete requests:      1000
Failed requests:        10
   (Connect: 0, Receive: 0, Length: 10, Exceptions: 0)
Non-2xx responses:      10
Total transferred:      275220 bytes
HTML transferred:       151300 bytes
Requests per second:    258.17 [#/sec] (mean)
Time per request:       774.689 [ms] (mean)
Time per request:       3.873 [ms] (mean, across all concurrent requests)
Transfer rate:          69.39 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       2
Processing:     2  427 890.3    136    3129
Waiting:        2  424 890.9    135    3128
Total:          3  427 890.3    137    3129

Percentage of the requests served within a certain time (ms)
  50%    137
  66%    181
  75%    213
  80%    231
  90%   3006
  95%   3079
  98%   3102
  99%   3111
 100%   3129 (longest request)
Finished 1000 requests
```

相比之下，database/sql 比 xorm 的响应速度快
