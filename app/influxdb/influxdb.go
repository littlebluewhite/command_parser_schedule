package influxdb

import (
	"command_parser_schedule/util/config"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"path/filepath"
	"runtime"
)

var (
	rootPath string
)

type Idb interface {
	Close()
	Writer() api.WriteAPIBlocking
	Querier() api.QueryAPI
}

type idb struct {
	client  influxdb2.Client
	writer  api.WriteAPIBlocking
	querier api.QueryAPI
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	rootPath = filepath.Dir(filepath.Dir(filepath.Dir(b)))
}

func NewInfluxdb(yamlName string) Idb {
	influxConfig := config.NewConfig[config.InfluxdbConfig](rootPath, "env", yamlName)
	dsn := fmt.Sprintf("http://%s:%s", influxConfig.Host, influxConfig.Port)
	client := influxdb2.NewClient(dsn, influxConfig.Token)
	writeAPI := client.WriteAPIBlocking(influxConfig.Org, influxConfig.Bucket)
	queryAPI := client.QueryAPI(influxConfig.Org)
	return &idb{
		client,
		writeAPI,
		queryAPI,
	}
}

func (i *idb) Close() {
	i.client.Close()
}

func (i *idb) Writer() api.WriteAPIBlocking {
	return i.writer
}

func (i *idb) Querier() api.QueryAPI {
	return i.querier
}
