## 测试用例

实现位置是`testcase/`

实现以下三个端：
* 客户端

  `testcase/client/client.go`
  ```bash
  # 启动命令
  make client
  ```

  请求上游服务

* 上游服务端

  `testcase/server/upstream.go`
  ```bash
  # 启动命令
  make up
  ```

  使用gin实现服务，并使用`pkg/hystrix/wrapper.go`中定义的中间件，对下游服务发起请求

* 下游服务端

  `testcase/server/downstream.go`
  ```bash
  # 启动命令
  make down
  ```

  初始化的时候传入一个成功率，当该服务被请求时，会根据定义的成功率返回请求成功(状态码200)或者失败(状态码500)
