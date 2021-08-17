package main

import "github.com/beego/beego/v2/server/web"

func main() {
	web.BConfig.Log.AccessLogs = true
	web.BConfig.WebConfig.DirectoryIndex = true
	web.SetStaticPath("/books", "books")
	web.SetStaticPath("/", "frontend/dist")
	web.SetStaticPath("/css", "frontend/dist/css")
	web.SetStaticPath("/img", "frontend/dist/img")
	web.SetStaticPath("/js", "frontend/dist/js")
	web.Run()
}
