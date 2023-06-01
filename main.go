package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

/*
手写web框架

	技术选型 基于 net/http标准库
		熟悉库
			库函数 > 结构定义 > 结构函数

			查询库的对外函数
				go doc net/http | grep "^func"

				去掉 设定的 New 和 Set 的设定属性的函数

				对剩下的函数进行分类
					为客户端提供创建HTTP服务的函数,Serve类似的
					为客户端提供调用的HTTP服务的类库,Get,Post类似的
					中转代理的函数,Proxy

				查看结构体
					go doc net/http | grep "^type"|grep struct
					通过go doc net/http.Server 查看结构体说明

				将要分析的代码从入口处一层层记录下来，每个函数，我们只记录其核心代码，然后对每个核心代码一层层解析。
*/
func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello,%q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
