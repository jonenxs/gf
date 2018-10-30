package main

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

//一个简单示例
func main() {
	s := g.Server()

	//:name、*any、{field}分别表示命名匹配规则、模糊匹配规则及字段匹配规则
	s.BindHandler("/:name", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})

	s.BindHandler("/:name/update", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})

	s.BindHandler("/:name/:action", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})

	s.BindHandler("/:name/*any", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})

	s.BindHandler("/user/list/{field}.html", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	//不同的规则中使用/符号来划分层级，层级越深的规则优先级也越高，gf框架的底层路由存储使用的是哈希表与数据链表构建的路由树
	s.SetPort(8199)
	s.Run()
}
