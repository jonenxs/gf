package main

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

func main() {
	s := g.Server()
	//[HTTPMethod:]路由规则[@域名]
	//Method：GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE）
	// 和@域名为非必需参数，一般来说直接给定路由规则参数即可，
	// BindHandler会自动绑定所有的请求方式，如果给定HTTPMethod，
	// 那么路由规则仅会在该请求方式下有效。@域名可以指定生效的域名名称，那么该路由规则仅会在该域名下生效
	//该路由仅在GET请求下有效
	s.BindHandler("GET:/{table}/list/{page}.html", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router)
	})
	//该路由规则仅会在GET请求及localhost域名下有效
	s.BindHandler("GET:/order/info/{order_id}@localhost", func(r *ghttp.Request) {
		r.Response.WriteJson(r.Router)
	})
	//该路由规则仅会在DELETE请求下有效
	s.BindHandler("DELETE:/comment/{id}", func(r *ghttp.Request) {
		r.Response.WriteJson(r.Router)
	})
	//其中返回的参数r.Router是当前匹配的路由规则信息，访问当该方法的时候，服务端会输出当前匹配的路由规则信息
	//在大多数场景下，我们很少直接在路由规则中使用@域名这样的规则来限定服务注册的域名，
	// 而是使用ghttp.Server.Domain(domains string)方法来获得指定域名列表的管理对象，
	// 随后使用该域名对象进行服务注册，域名对象即可实现对指定域名的绑定操作
	s.SetPort(8199)
	s.Run()
}
