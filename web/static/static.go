package main

import "gitee.com/johng/gf/g"

//静态服务
func main() {
	s := g.Server()
	//创建了Web Server对象之后，我们可以使用Set*方法来设置Web Server的属性
	//SetIndexFolder用来设置是否允许列出Web Server主目录的文件列表（默认为false）；
	s.SetIndexFolder(true)
	//SetServerRoot用来设置Web Server的主目录（类似其他Web Server中的DocumentRoot地址，默认为空，
	// 在某些时候，Web Server仅提供接口服务，因此Web Server的主目录为非必需参数）
	//Web Server默认情况下是没有任何主目录的设置，只有设置了主目录，才支持对应主目录下的静态文件的访问
	s.SetServerRoot("/home/www/")
	s.Run()
}
