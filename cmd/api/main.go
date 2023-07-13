package main

import (
	"command_parser_schedule/app/dbs"
	_ "command_parser_schedule/docs"
	"command_parser_schedule/gin/group"
	"command_parser_schedule/gin/initial"
	"command_parser_schedule/util/config"
	"command_parser_schedule/util/logFile"
	"context"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	mainLog logFile.LogFile
)

// 初始化配置
func init() {
	// log配置
	mainLog = logFile.NewLogFile("", "main.log")
}

// @title           Schedule-Command swagger API
// @version         2.0
// @description     This is a schedule-command server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Wilson
// @contact.url    https://github.com/littlebluewhite
// @contact.email  wwilson008@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:5487

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mainLog.Info().Println("command module start")

	// DBs start includes SQL Cache
	DBS := dbs.NewDbs(mainLog, false)
	defer func() {
		DBS.GetIdb().Close()
		mainLog.Info().Println("influxDB Disconnect")
	}()

	// gin app start
	ginApp := initial.NewGinApp(mainLog, DBS)

	// injection
	group.Inject(ginApp)

	// server config
	ServerConfig := config.NewConfig[config.ServerConfig](".", "env", "server")

	// server
	var sb strings.Builder
	sb.WriteString(":")
	sb.WriteString(ServerConfig.Port)
	srv := &http.Server{
		Addr:           sb.String(),
		Handler:        ginApp.GetRouter(),
		ReadTimeout:    ServerConfig.ReadTimeout * time.Second,
		WriteTimeout:   ServerConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// API server Start
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mainLog.Error().Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	mainLog.Info().Println("API server shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		mainLog.Error().Fatal("Server forced to shutdown: ", err)
	}

	mainLog.Info().Println("Server exiting")

}
