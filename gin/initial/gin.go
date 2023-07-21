package initial

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/app/time_server"
	"command_parser_schedule/gin/middleware"
	"command_parser_schedule/util/logFile"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type GinApp interface {
	GetRouter() *gin.Engine
	GetDbs() dbs.Dbs
	GetTimeServer() time_server.TimeServer
}

type ginApp struct {
	Router *gin.Engine
	Dbs    dbs.Dbs
	ts     time_server.TimeServer
}

func NewGinApp(log logFile.LogFile, dbs dbs.Dbs, ts time_server.TimeServer) GinApp {
	r := initRouter(log)
	return &ginApp{
		Router: r,
		Dbs:    dbs,
		ts:     ts,
	}
}

func initRouter(log logFile.LogFile) *gin.Engine {
	ginFile, err := os.OpenFile("./log/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error().Fatal("can not open log file: " + err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(ginFile, os.Stdout)
	r := gin.Default()

	// cors middleware
	r.Use(middleware.CORSMiddleware())
	gin.SetMode(gin.ReleaseMode)
	return r
}
func (g *ginApp) GetRouter() *gin.Engine {
	return g.Router
}

func (g *ginApp) GetDbs() dbs.Dbs {
	return g.Dbs
}

func (g *ginApp) GetTimeServer() time_server.TimeServer {
	return g.ts
}
