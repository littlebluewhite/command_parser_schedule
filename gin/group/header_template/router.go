package header_template

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
		panic("initial header template router error")
	}
	h := Handler{
		O: r.O,
		L: r.L,
	}

	// set api group
	g := r.R.Group("/header_template")

	g.GET("/", h.GetheaderTemplates)
	g.GET("/:id", h.GetHeaderTemplateById)
	g.POST("/", h.AddHeaderTemplate)
	g.PATCH("/", h.UpdateHeaderTemplate)
	g.DELETE("/", h.DeleteHeaderTemplate)
}
