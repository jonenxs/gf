package main

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

//动态路由规则
func main() {
	s := g.Server()

	//动态路由规则分为三种：命名匹配规则、模糊匹配规则和字段匹配规则
	//一个简单的分页路由实例
	s.BindHandler("/user/list/{page}.html", func(r *ghttp.Request) {
		r.Response.Writeln(r.Get("page"))
	})
	//{***}规则与 :*** 规则混合使用
	s.BindHandler("/{object}/:attr/{act}.php", func(r *ghttp.Request) {
		r.Response.Writeln(r.Get("object"))
		r.Response.Writeln(r.Get("attr"))
		r.Response.Writeln(r.Get("act"))
	})
	// 多种模糊匹配规则混合使用
	s.BindHandler("/{class}-{course}/:name/*act", func(r *ghttp.Request) {
		r.Response.Writeln("class")
		r.Response.Writeln("course")
		r.Response.Writeln("name")
		r.Response.Writeln("act")
	})
	s.SetPort(8199)
	s.Run()
}
