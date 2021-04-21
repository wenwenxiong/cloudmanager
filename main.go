package main

import (
	_ "cloudmanager/boot"
	_ "cloudmanager/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
