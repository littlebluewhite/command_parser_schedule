package sql

import (
	"command_parser_schedule/util/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	rootPath string
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	rootPath = filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(b))))
}

func NewDB(dirPath, fileName, yamlName string) (*gorm.DB, error) {
	newPath := filepath.Join("./log", dirPath)
	_ = os.MkdirAll(newPath, os.ModePerm)
	newPath = filepath.Join(newPath, fileName)
	file, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("can not open log file: " + err.Error())
	}

	newLogger := logger.New(
		log.New(io.MultiWriter(file, os.Stdout), "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	dbConfig := config.NewConfig[config.DBConfig](rootPath, "env", yamlName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&parseTime=true&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DB)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
		PrepareStmt:            true,
	})
}
