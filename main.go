package main

import (
	"ayachanV2/Config"
	"ayachanV2/Controllers"
	"ayachanV2/Databases"
	"ayachanV2/Router"
	"log"
	"time"
)

// @title ayachan API
// @version 2.0
// @description api 计算Bestdori谱面难度，获得Bestdori数据，常见Bandori谱面格式转换等

// @contact.name 6QHTSK

// @license.name MIT
// @license.url https://mit-license.org/

// @host 127.0.0.1:8080
// @BasePath /v2

func main() {
	defer Databases.SqlDB.Close()
	Config.InitConfig()
	var lastUpdate time.Time
	lastUpdate, err := Databases.GetLastUpdate()
	if err != nil {
		//log.Fatalln(err.Error())
		log.Println("读表失败，表为空，最后更新设为0")
	}
	Config.SetLastUpdate(lastUpdate)

	//errCode, err := Services.BestdoriFanMadeSyncAll()
	//if err != nil {
	//	log.Fatalln(errCode, err.Error())
	//}

	Controllers.CronSync()

	router := Router.InitRouter()

	//Router.InitSwaggerDoc(router)
	Router.InitAPIV2(router)
	//
	_ = router.Run("0.0.0.0:8080")
}
