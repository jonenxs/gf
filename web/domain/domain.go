package main

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

func Hello1(r *ghttp.Request) {
	r.Response.Write("127.0.0.1: Hello1!")
}

func Hello2(r *ghttp.Request) {
	r.Response.Write("127.0.0.1: Hello2!")
}

//域名&多域名
func main() {
	s := g.Server()
	//我们访问http://127.0.0.1/和http://localhost/可以看输出不同的内容
	s.Domain("127.0.0.1").BindHandler("/", Hello1)
	s.Domain("localhost").BindHandler("/", Hello2)
	//Domain方法支持多个域名参数，使用英文“,”号分隔
	//s.Domain("localhost1,localhost2,localhost3").BindHandler("/", Hello2)
	//Domain方法的参数必须是准确的域名，不支持泛域名形式，例如：*.johng.cn或者.johng.cn是不支持的，api.johng.cn或者johng.cn才被认为是正确的域名参数。
	s.Run()
}
