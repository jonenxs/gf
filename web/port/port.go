package main

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

//多端口监听
func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writefln("go frame!")
	})
	//ghttp.Server同时支持多端口监听，只需要往SetPort参数设置绑定多个端口号即可
	//（当然，针对于HTTPS服务，我们也同样可以通过SetHTTPSPort来设置绑定并支持多个端口号的监听，HTTPS服务的介绍请参看后续对应章节）
	s.SetPort(8100,8200,8300)
	s.Run()
}
