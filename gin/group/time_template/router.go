package time_template

import (
	"command_parser_schedule/util/logFile"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	R *gin.Engine
	O Operate
	L logFile.LogFile
}

func InitRoutes(r Routes) {
	err := r.O.ReloadCache()
	if err != nil {
		panic("initial time template router error")
	}
	h := Handler{
		O: r.O,
		L: r.L,
	}

	// set api group
	g := r.R.Group("/time_template")

	g.GET("/", h.GetTimeTemplates)
	g.GET("/:id", h.GetTimeTemplateById)
	g.POST("/", h.AddTimeTemplate)
	g.PATCH("/", h.UpdateTimeTemplate)
	g.DELETE("/", h.DeleteTimeTemplate)
}
