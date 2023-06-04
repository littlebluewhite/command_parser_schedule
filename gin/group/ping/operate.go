package ping

import (
	"command_parser_schedule/gin/initial"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type Operate interface {
}

type operate struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewOperate(dbs initial.Dbs) Operate {
	return &operate{
		db:    dbs.GetSql(),
		cache: dbs.GetCache(),
	}
}
