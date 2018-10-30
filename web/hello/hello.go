package main

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)
//哈喽世界
func main() {
	//通过g.Server()方法获得一个默认的Web Server对象
	//该方法采用单例模式设计，也就是说，多次调用该方法，返回的是同一个Web Server对象
	// （我们也可以通过ghttp模块的ghttp.GetServer()来获取单例的Web Server对象
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
	})
	//通过Run()方法执行Web Server的监听运行，在没有任何额外设置的情况下，它默认监听80端口
	s.Run()
}
