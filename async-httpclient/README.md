## HTTP Reactive Client 
这是一个典型的消息（事件）驱动的案例。

#### 1. 依据文档图6-1，用中文描述 Reactive 动机

同步方式的性能差而且浪费资源；
reactive编程提供异步方式，而且为了限制流量并实现更低的延迟，通过单独一个响应将所有必要的信息返回给客户端。
图6-1
![image 6-1](https://github.com/karl-jm-huang/golang-learning/blob/master/async-httpclient/images/6-1.png)

描述：该层接受来自外部的请求，并负责调用对内部服务的多个请求。当来自内部服务的响应在该层中可用时，它们被组合成单个响应，并被发送回客户端。

#### 2. 使用 go HTTPClient 实现图 6-2 的 Naive Approach
图6-2
![image 6-2](https://github.com/karl-jm-huang/golang-learning/blob/master/async-httpclient/images/6-2.png)

#### 3. 为每个 HTTP 请求设计一个 goroutine ，利用 Channel搭建基于消息的异步机制，实现图 6-3
图6-3
![image 6-3](https://github.com/karl-jm-huang/golang-learning/blob/master/async-httpclient/images/6-3.png)

#### 4. 对比两种实现，用数据说明 go 异步 REST 服务协作的优势

从响应时间可以看出 go 异步 REST 服务协作能缩短用户等待时间
