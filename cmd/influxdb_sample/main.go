package main

import (
	"command_parser_schedule/app/dbs/influxdb"
	"command_parser_schedule/dal/model"
	"context"
	"fmt"
	"github.com/goccy/go-json"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main() {
	d := model.CommandTemplate{ID: 1, Name: "aaa"}
	j, _ := json.Marshal(d)
	ctx := context.Background()
	idb := influxdb.NewInfluxdb("influxdb")
	defer idb.Close()
	p := influxdb2.NewPoint("schedule_history",
		map[string]string{"id": "2", "name": "alarm SOP", "user": "wilson"},
		map[string]interface{}{"data": j},
		time.Now())
	p2 := influxdb2.NewPoint("schedule_history",
		map[string]string{"id": "1", "name": "alarm SOP", "user": "wilson"},
		map[string]interface{}{"complete": 1, "duration": 2},
		time.Now())
	if err := idb.Writer().WritePoint(ctx, p); err != nil {
		panic(err)
	}

	if err := idb.Writer().WritePoint(ctx, p2); err != nil {
		panic(err)
	}

	result, err := idb.Querier().Query(ctx, `from(bucket:"schedule")
|> range(start: -2h)
|> filter(fn: (r) => r._measurement == "schedule_history")`)
	if err == nil {
		for result.Next() {
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			v := result.Record().Value()
			var c model.CommandTemplate
			if result.Record().Field() == "data" {
				json.Unmarshal([]byte(v.(string)), &c)
				fmt.Println(c)
			}
			fmt.Printf("value: %v\ntype: %T\n", v, v)
			fmt.Printf("values: %v\n", result.Record().Values())
			fmt.Printf("result: %v\n", result.Record().Result())
			fmt.Printf("measurement: %v\n", result.Record().Measurement())
			fmt.Printf("field: %v\n", result.Record().Field())
			fmt.Printf("table: %v\n", result.Record().Table())
			fmt.Printf("start: %v\n", result.Record().Start())
			fmt.Printf("stop: %v\n", result.Record().Stop())
			fmt.Printf("time: %v\n", result.Record().Time())
			fmt.Printf("value by key(id): %v\n", result.Record().ValueByKey("id"))
		}
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}
