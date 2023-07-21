package schedule

import (
	"command_parser_schedule/gin/initial"
	"command_parser_schedule/util/logFile"
)

func Inject(ginApp initial.GinApp) {
	o := NewOperate(ginApp.GetDbs(), ginApp.GetTimeServer())
	log := logFile.NewLogFile("router", "schedule.log")
	InitRoutes(Routes{
		ginApp.GetRouter(),
		o,
		log,
	})

}
