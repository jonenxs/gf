package main

import (
	"gitee.com/johng/gf/g"
)

//多服务支持
func main() {
	//在支持多个Web Server的语句中，给g.Server方法传递了不同的参数
	//（参数可以为任意类型，常用字符串或者整型识别），该参数为Web Server的名称，
	// 之前我们提到g.Server方法采用了单例设计模式，该参数用于标识不同的Web Server，因此需要保证唯一性
	s1 := g.Server("s1")
	s1.SetPort(8080)
	s1.SetIndexFolder(true)
	s1.SetServerRoot("/home/www/static1")
	s1.Start()

	s2 := g.Server("s2")
	s2.SetPort(8080)
	s2.SetIndexFolder(true)
	s2.SetServerRoot("/home/www/static2")
	s2.Start()
	//如果需要获取同一个Web Server，那么传入同一个名称即可。
	// 例如在多个goroutine中，或者不同的模块中，都可以通过g.Server获取到同一个Web Server对象。
}
