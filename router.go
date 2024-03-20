package main

import (
	"github.com/6QHTSK/ayachan/config"
	"github.com/6QHTSK/ayachan/controller"
	_ "github.com/6QHTSK/ayachan/docs"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func InitRouter() (app *iris.Application) {
	app = iris.New()
	if config.Config.Debug {
		app.Logger().SetLevel("debug")
	} else {
		app.Logger().SetLevel("error")
	}
	app.Use(recover.New())
	app.Use(logger.New())
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	return app
}

func InitAPI(app *iris.Application) {
	v2 := app.Party("/v2")
	{
		v2.Get("/version", controller.GetVersion)
		chartInfo := v2.Party("/chart/metrics")
		{
			chartInfo.Get("/bestdori/{chartID:int}", controller.ChartMetricsFromBestdori)
			chartInfo.Get("/bandori/{chartID:int}/{diffStr}", controller.ChartMetricsFromBandori)
			chartInfo.Post("/custom/{diffStr}", controller.ChartMetrics)
		}
		// swagger
		swaggerUI := swagger.CustomWrapHandler(&swagger.Config{
			// The url pointing to API definition.
			URL:         "doc.json",
			DeepLinking: true,
		}, swaggerFiles.Handler)

		// Register on http://localhost:8080/swagger
		v2.Get("/doc", swaggerUI)
		// And the wildcard one for index.html, *.js, *.css and e.t.c.
		v2.Get("/doc/{any:path}", swaggerUI)
	}
}
