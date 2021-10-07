# idefav-httpserver
## 原理
使用 DispatchHandler 统一处理请求
![image](https://user-images.githubusercontent.com/6405415/136300504-2d4e0179-3366-4207-b534-ea3ceb8aecbc.png)

## 如何编写 Handler
1. 在 handler 文件夹添加 handler 文件夹, 然后添加 demo.go 和 init.go 文件
2. 在 demo.go 文件添加 DemoHandler struct
3. DemoHandler 实现 HandlerMapping 接口
HandlerMapping 包含的函数说明:
* `Name()` 返回 Handler 名称
* `Path()` 返回 Handler 匹配路径
* `Method()` 返回 Handler 可以处理的 HttpMethod
* `Handler()` 开发真正的处理逻辑
4. 在 init.go 编写 包初始化函数
```go
func init(){
 // init 
   headerz:=HeaderHandler{}
   handler.DefaultDispatchHandler.AddHandler(&headerz)
}
```
5. 在 server.go 新增包引用
```go
import (
    //...
    _ "idefav-httpserver/handler/headerz"
)
```
6. 启动服务

