package initial

import (
	"command_parser_schedule/app/sql"
	"command_parser_schedule/util/logFile"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"time"
)

type Dbs interface {
	initSql(log logFile.LogFile)
	initCache()
	GetSql() *gorm.DB
	GetCache() *cache.Cache
}

type dbs struct {
	Sql   *gorm.DB
	Cache *cache.Cache
}

func NewDbs(log logFile.LogFile) Dbs {
	d := &dbs{}
	d.initSql(log)
	d.initCache()
	return d
}

// DB start
func (d *dbs) initSql(log logFile.LogFile) {
	s, err := sql.NewDB("mySQL", "DB.log", "db")
	if err != nil {
		log.Error().Println("DB Connection failed")
		panic(err)
	} else {
		log.Info().Println("DB Connection successful")
	}
	d.Sql = s
}

func (d *dbs) initCache() {
	d.Cache = cache.New(5*time.Minute, 10*time.Minute)
}

func (d *dbs) GetSql() *gorm.DB {
	return d.Sql
}

func (d *dbs) GetCache() *cache.Cache {
	return d.Cache
}
