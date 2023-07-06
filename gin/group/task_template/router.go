package task_template

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
	g := r.R.Group("/task_template")

	g.GET("/", h.GetTaskTemplates)
	g.GET("/:id", h.GetTaskTemplateById)
	g.POST("/", h.AddTaskTemplate)
	g.DELETE("/", h.DeleteTaskTemplate)
}
