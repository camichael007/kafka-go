package main

import (
	_ "kafka/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

