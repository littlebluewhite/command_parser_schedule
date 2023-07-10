package schedule

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
	g := r.R.Group("/schedule")

	g.GET("/", h.GetSchedules)
	g.GET("/:id", h.GetScheduleById)
	g.POST("/", h.AddSchedule)
	g.PATCH("/", h.UpdateSchedule)
	g.DELETE("/", h.DeleteSchedule)
}
