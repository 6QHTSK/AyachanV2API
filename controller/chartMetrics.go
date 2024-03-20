package controller

import (
	"github.com/6QHTSK/ayachan/model/bestdoriChart"
	"github.com/6QHTSK/ayachan/service"
	"github.com/kataras/iris/v12"
)

func convertDiffStr(diffStr string) (diff int) {
	var DiffStrDict = map[string]int{
		"easy": 0, "ez": 0,
		"normal": 1, "nm": 1,
		"hard": 2, "hd": 2,
		"expert": 3, "ex": 3,
		"special": 4, "sp": 4,
	}
	diff, ok := DiffStrDict[diffStr]
	if !ok {
		return 3
	}
	return diff
}

type chartMetricsFromBandoriParam struct {
	ChartID int    `param:"chartID" validate:"required"`
	DiffStr string `param:"diffStr" validate:"required"`
}

// ChartMetricsFromBandori 计算Bandori谱面的各项信息
// @Version 2.2
// @Description.markdown
// @Summary 计算Bandori谱面的各项信息
// @Tags ChartMetrics
// @Produce json
// @Param chartID path int true "谱面ID"
// @Param diffStr path string true "难度字符串，建议在[easy,normal,hard,expert,special]中选择"
// @Success 200 {object} model.ChartMetrics "Description中所提及的各项信息"
// @Failure 400 {object} ErrorResponse "传入Param错误"
// @Failure 500 {object} ErrorResponse "服务器内部错误，包括找不到谱面等"
// @Router /v2/chart/metrics/bandori/{chartID}/{diffStr} [get]
func ChartMetricsFromBandori(ctx iris.Context) {
	var params chartMetricsFromBandoriParam
	if err := ctx.ReadParams(&params); err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, failedResponse(validationErrorParser(err)))
		return
	}
	MapInfo, err := service.BandoriChartAnalyze(params.ChartID, convertDiffStr(params.DiffStr))
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusInternalServerError, failedResponse(err.Error()))
		return
	}
	_ = ctx.JSON(MapInfo)
}

type chartMetricsFromBestdoriParam struct {
	ChartID int `param:"chartID" validate:"required"`
}

// ChartMetricsFromBestdori 计算Bandori谱面的各项信息
// @Version 2.2
// @Description.markdown
// @Summary 计算Bestdori谱面的各项信息，谱面的难度将会根据Bestdori上谱面声称的难度进行选择
// @Tags ChartMetrics
// @Produce json
// @Param chartID path int true "谱面ID"
// @Success 200 {object} model.ChartMetrics "Description中所提及的各项信息"
// @Failure 400 {object} ErrorResponse "传入Param错误"
// @Failure 500 {object} ErrorResponse "服务器内部错误，包括找不到谱面等"
// @Router /v2/chart/metrics/bestdori/{chartID} [get]
func ChartMetricsFromBestdori(ctx iris.Context) {
	var params chartMetricsFromBestdoriParam
	if err := ctx.ReadParams(&params); err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, failedResponse(validationErrorParser(err)))
		return
	}
	MapInfo, err := service.BestdoriChartAnalyze(params.ChartID)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusInternalServerError, failedResponse(err.Error()))
		return
	}
	_ = ctx.JSON(MapInfo)
}

type chartMetricsParams struct {
	DiffStr string `param:"diffStr" validate:"required"`
}

// ChartMetrics 计算上传谱面的各项信息
// @Version 2.2
// @Summary 计算上传谱面的各项信息
// @Description.markdown
// @Tags ChartMetrics
// @Accept json
// @Produce json
// @Param diffStr path string true "难度字符串，建议在[easy,normal,hard,expert,special]中选择择"
// @Param message body bestdoriChart.BestdoriV2Chart true "BestdoriV2谱面"
// @Success 200 {object} model.ChartMetrics "Description中所提及的各项信息"
// @Failure 400 {object} ErrorResponse "传入谱面/Param错误"
// @Failure 500 {object} ErrorResponse "服务器内部错误，包括找不到谱面等"
// @Router /v2/chart/metrics/custom/{diffStr} [post]
func ChartMetrics(ctx iris.Context) {
	var chart bestdoriChart.BestdoriV2Chart
	var params chartMetricsParams

	if err := ctx.ReadJSON(&chart); err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, failedResponse("无法处理传入谱面"))
		return
	}

	if err := ctx.ReadParams(&params); err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, failedResponse("无法解析diff"))
		return
	}

	MapInfo, err := service.ChartAnalyze(&chart, convertDiffStr(params.DiffStr))
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, failedResponse(err.Error()))
		return
	}

	_ = ctx.JSON(MapInfo)
}
