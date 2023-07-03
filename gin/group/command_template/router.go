package command_template

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
	h := Handler{
		O: r.O,
		L: r.L,
	}

	// set api group
	g := r.R.Group("/command_template")

	g.GET("/", h.GetcommandTemplates)
	g.GET("/:id", h.GetCommandTemplateById)
	g.POST("/", h.AddCommandTemplate)
	g.PATCH("/", h.UpdateCommandTemplate)
	g.DELETE("/", h.DeleteCommandTemplate)
}
