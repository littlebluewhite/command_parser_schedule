package ping

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
	g := r.R.Group("/ping")

	g.GET("/test", h.GetPing)
	g.GET("/list", h.GetListPing)
}
