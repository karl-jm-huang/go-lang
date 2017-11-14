### 使用Martini框架实现小server
### 测试
运行server，设置监听端口3000，不设置则监听默认端口8080
```
~$ go run main.go -p=3000
[martini] listening on :3000 (development)
```

### curl
另开一个终端执行curl，可以看到返回用户HuangJM已经登录的信息
```
~$ curl -v http://localhost:3000/login/HuangJM
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 3000 (#0)
> GET /login/HuangJM HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 14 Nov 2017 14:40:19 GMT
< Content-Length: 28
< Content-Type: text/plain; charset=utf-8
<
user HuangJM has logged in.
* Connection #0 to host localhost left intact

```

### ab

```
~$ ab -n 1000 -c 100 http://localhost:3000/login/HuangJM
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
Server Port:            3000

Document Path:          /login/HuangJM
Document Length:        28 bytes

Concurrency Level:      100
Time taken for tests:   0.071 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      145000 bytes
HTML transferred:       28000 bytes
Requests per second:    14094.04 [#/sec] (mean)
Time per request:       7.095 [ms] (mean)
Time per request:       0.071 [ms] (mean, across all concurrent requests)
Transfer rate:          1995.74 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.9      1       4
Processing:     1    6   3.0      6      13
Waiting:        1    5   3.1      6      13
Total:          2    7   2.6      7      15

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      8
  75%      8
  80%      8
  90%     10
  95%     13
  98%     14
  99%     14
 100%     15 (longest request)

```

从结果看到，我们服务器对1000个 请求全部响应成功，每秒请求量14094.04，平均请求等待时间7.095ms，传输速率1995.74Kb/s
