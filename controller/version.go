package controller

import (
	"github.com/6QHTSK/ayachan/config"
	"github.com/kataras/iris/v12"
)

type APIVersion struct {
	Version string `json:"version"`
}

// GetVersion 获得该API的版本
// @version 2.2
// @Description 根据内部信息得到API的版本
// @Summary 获得API版本
// @Tags Version
// @Produce plain
// @Success 200 {object} APIVersion "获得的API版本号"
// @Router /v2/version [get]
func GetVersion(ctx iris.Context) {
	_ = ctx.JSON(APIVersion{Version: config.Version})
}
