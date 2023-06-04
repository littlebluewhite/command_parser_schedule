package ping

import (
	"command_parser_schedule/gin/initial"
	"command_parser_schedule/util/logFile"
)

func Inject(ginApp initial.GinApp) {
	o := NewOperate(ginApp.GetDbs())
	log := logFile.NewLogFile("router", "ping.log")
	InitRoutes(Routes{
		ginApp.GetRouter(),
		o,
		log,
	})

}
