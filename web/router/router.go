package main

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

//强大路由特性
func main() {
	s := g.Server()
	//这是一个混合的路由规则示例，用于展示某个班级、某个学科、某个学生、对应的操作
	s.BindHandler("/{class}-{course}/:name/*act", func(r *ghttp.Request) {
		r.Response.Writeln(r.Get("class"))
		r.Response.Writeln(r.Get("course"))
		r.Response.Writeln(r.Get("name"))
		r.Response.Writeln(r.Get("act"))
	})
	s.SetPort(8199)
	s.Run()

}
