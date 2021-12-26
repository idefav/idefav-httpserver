# idefav-httpserver

## 原理

使用 DispatchHandler 统一处理请求
![image](https://user-images.githubusercontent.com/6405415/136300504-2d4e0179-3366-4207-b534-ea3ceb8aecbc.png)

## 项目介绍

1. 项目整体采用模块化结构, 易扩展
2. 多级配置系统 配置优先级: cmd > env > configFile 支持 yaml 文件配置
3. 加入预热模块, 支持启动预热, 在预热期间, /healthz 接口将返回 UnHealthy 状态
4. 加入优雅关机模块, 支持系统结束到 CTL+C 或者 kill命令时, 支持 clean 相关资源, /healthz 接口会变成 UnHealthy 状态, 和预热功能结合, 再用 K8s 的readinessProbe
   可以实现真正的流量无损发布
5. 加入 Router 模块, 默认路由不支持复杂 Url 匹配, 但是路由模块支持扩展
6. 新增拦截器模块, 支持对请求进行自定义过滤类需求, 比如加入随机延迟
7. 新增metrics支持, 支持暴露metrics指标给prometheus
8. 新增opentracing支持, 支持istio envoy trace id透传

## 如何编写 Handler

1. 在 handler 文件夹添加 逻辑处理 文件夹, 然后添加 demo.go 文件
2. demo.go 示例代码如下

```go
package demo

import (
	"idefav-httpserver/cfg"
	"idefav-httpserver/handler"
	"net/http"
	"os"
)

func init() {
	handler.DefaultDispatchHandler.AddHandler(&handler.SimpleHandler{
		Name:   "Headerz",
		Path:   "/headerz",
		Method: http.MethodGet,
		Proc: func(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
			for headerName, headerValues := range request.Header {
				for _, v := range headerValues {
					writer.Header().Add(headerName, v)
				}
			}
			version := os.Getenv(cfg.VERSION)
			if version != "" {
				writer.Header().Add(cfg.VERSION, version)
			}
			return "Ok", nil
		},
	})

	handler.DefaultDispatchHandler.AddHandler(&handler.SimpleHandler{
		Name:   "Demo",
		Path:   "/demo",
		Method: http.MethodGet,
		Proc: func(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
			return "Demo", nil
		},
	})

	handler.DefaultDispatchHandler.AddHandler(&handler.SimpleHandler{
		Name:   "Panic",
		Path:   "/panic",
		Method: http.MethodGet,
		Proc: func(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
			panic("demo panic")
		},
	})
}


```

5. 在 auto/handler.go 新增包引用

```go
import (
// ...
_ "idefav-httpserver/handler/demo"
)
```

6. 启动服务
访问 http://localhost:8081/healthz 
返回结果: 
```text
{
"code": 0,
"message": "",
"data": "Health"
}
```

## 如何自定义预热逻辑

1. 在 components/warmup/components 新增 demo.go
2. demo.go 内容如下
```go
package components

import (
	"idefav-httpserver/components/warmup"
	"log"
)

func init() {
	warmup.Add("Demo", func() error {
		log.Println("demo warmup!")
		return nil
	})
}
```
3. 启动项目, 查看日志, 是否有 `demo warmup!` 的日志打出
```text
2021/11/28 20:30:13 Server Started
2021/11/28 20:30:13 Warmup now!
2021/11/28 20:30:13 Warmup completed!
2021/11/28 20:30:13 Server listen at::8081
2021/11/28 20:30:13 demo warmup!
```

## 如何自定义优雅停机逻辑
1. 在 /components/shutdown/components 目录下 新增 demo.go
2. 代码如下
```go
package components

import (
	"idefav-httpserver/components/shutdown"
	"log"
	"math"
	"time"
)

func init() {
	shutdown.Add(&shutdown.DefaultGracefulShutdownComponent{
		Name:  "Demo",
		Order: math.MinInt,
		Proc: func() {
			log.Println("Cleaning...")
			time.Sleep(2 * time.Second)
			log.Println("Clean Done!")
		},
	})
}

```
3. 启动项目, 然后停止项目, 查看是否出现相关日志
```text
021/11/28 20:32:23 Server is shutting down and begin cleaning!
2021/11/28 20:32:23 Graceful component executing:Demo
2021/11/28 20:32:23 Cleaning...
2021/11/28 20:32:25 Clean Done!
2021/11/28 20:32:25 Server is down, and clean completed!
```

## 如何扩展 router 组件
实现 router.Interface 接口
```go
type Interface interface {
	GetName() string
	Add(handler models.HandlerMapping)
	Match(request *http.Request) (models.HandlerMapping, error)
}
```

## opentracing 支持, 支持 Istio trace id透传
```go
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Span    opentracing.Span
}
```
使用参照httpbin接口调用demo
```go
asyncReq, err := http.NewRequest(http.MethodGet, "http://httpbin/headers", nil)
err = tracing.Inject(ctx.Span, asyncReq)
if err != nil {
    ctx.Span.SetTag("error", true)
    ctx.Span.LogEvent(fmt.Sprintf("Could not inject span context into header: %v", err))
}
resp, err := http.DefaultClient.Do(asyncReq)
if err != nil {
    ctx.Span.SetTag("error", true)
    ctx.Span.LogEvent(fmt.Sprintf("%s /headers error: %v", http.MethodGet, err))
    return nil, err
}
if resp.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("access httpbin failed, %v", resp.StatusCode)
}
body, err := ioutil.ReadAll(resp.Body)
var data = make(map[string]interface{})
err = json.Unmarshal(body, &data)
return data, err
```