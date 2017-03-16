package main

import (
	_ "webtest/routers"
	"github.com/astaxie/beego"


)
func main() {
	beego.SetStaticPath("/upload", "upload")
	beego.Run()
}

