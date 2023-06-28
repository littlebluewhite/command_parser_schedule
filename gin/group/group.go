package group

import (
	"command_parser_schedule/gin/group/ping"
	"command_parser_schedule/gin/group/time_template"
	"command_parser_schedule/gin/initial"
	"command_parser_schedule/gin/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Inject(ginApp initial.GinApp) {
	//get router from ginApp
	r := ginApp.GetRouter()

	// middleware
	r.Use(middleware.Latency())

	// inject routers
	ping.Inject(ginApp)
	time_template.Inject(ginApp)

	// swagger router
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
