package main

import (
	"command_parser_schedule/app/dbs/rdb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	r := rdb.NewRedis("redis")
	ctx := context.Background()
	if _, err := r.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, "key", "str1", "hello")
		rdb.HSet(ctx, "key", "str2", "world")
		rdb.HSet(ctx, "key", "int", "123")
		rdb.HSet(ctx, "key", "bool", "1")
		return nil
	}); err != nil {
		panic(err)
	}
	if err := r.HDel(ctx, "key", "bool").Err(); err != nil {
		panic(err)
	}

	jsonData, _ := json.Marshal(map[string]string{"vd": "dd"})

	setData := make(map[string]interface{})
	setData["asdf"] = 44
	setData["zxve"] = "aaabbbd"
	setData["zxve"] = string(jsonData)

	if err := r.HMSet(ctx, "key1", setData).Err(); err != nil {
		panic(err)
	}

	data, err := r.HGetAll(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	data3, err := r.HGetAll(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}

	// Scan a subset of the fields.
	data2, err := r.HMGet(ctx, "key", "str1", "int").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(data2)
	fmt.Println(data3)
}
