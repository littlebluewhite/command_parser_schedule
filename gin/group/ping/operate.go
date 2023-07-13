package ping

import (
	"command_parser_schedule/app/dbs"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type Operate interface {
}

type operate struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewOperate(dbs dbs.Dbs) Operate {
	return &operate{
		db:    dbs.GetSql(),
		cache: dbs.GetCache(),
	}
}
