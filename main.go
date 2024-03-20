package main

import (
	"flag"
	"fmt"
	"github.com/6QHTSK/ayachan/config"
)

var showVer bool
var runAddr string

func init() {
	flag.BoolVar(&showVer, "v", false, "查看版本号")
	flag.StringVar(&runAddr, "a", config.Config.RunAddr, "运行地址")
}

// main Ayachan 谱面难度分析器
// @Title	Ayachan Bandori谱面难度分析器
// @version		2.2
// @description	可对谱面特征进行提取，并拟合谱面难度难度模块
// @description.markdown
// @license.name	MIT
// @BasePath	/v2
func main() {
	flag.Parse()
	if showVer {
		fmt.Println(config.Version)
	} else {
		router := InitRouter()

		InitAPI(router)
		_ = router.Listen(runAddr)
	}
}
