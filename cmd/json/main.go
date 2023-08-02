package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data, _ := json.Marshal([]map[string]interface {
	}{{
		"a": []any{1, "5"},
		"b": "B",
		"c": 1,
		"d": "2",
	}, {
		"e": 2,
	},
	})
	fmt.Printf("%T, %+v\n", data, data)
	var data2 []map[string]interface{}
	e := json.Unmarshal(data, &data2)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("%T, %+v\n", data2, data2)
	for _, item := range data2 {
		for _, v := range item {
			switch v.(type) {
			case float64:
				fmt.Printf("float64: %v\n", v)
			case string:
				fmt.Printf("string: %s\n", v)
			default:
				fmt.Printf("%T, %v\n", v, v)
			}
		}
	}

	for _, v := range data2[0]["a"].([]any) {
		switch v.(type) {
		case float64:
			fmt.Printf("float64: %v\n", v)
		case string:
			fmt.Printf("string: %s\n", v)
		default:
			fmt.Printf("%T, %v\n", v, v)
		}
	}

	data3, _ := json.Marshal(1)
	var data4 any
	json.Unmarshal(data3, &data4)
	switch data4.(type) {
	case float64:
		fmt.Printf("float64: %v\n", data4)
	case string:
		fmt.Printf("string: %s\n", data4)
	default:
		fmt.Printf("%T, %v\n", data4, data4)
	}
}
