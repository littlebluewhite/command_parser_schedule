package initial

import (
	"command_parser_schedule/gin/middleware"
	"command_parser_schedule/util/logFile"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type GinApp interface {
	GetRouter() *gin.Engine
	GetDbs() Dbs
}

type ginApp struct {
	Router *gin.Engine
	Dbs    Dbs
}

func NewGinApp(log logFile.LogFile, dbs Dbs) GinApp {
	r := initRouter(log)
	return &ginApp{
		Router: r,
		Dbs:    dbs,
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

func (g *ginApp) GetDbs() Dbs {
	return g.Dbs
}
