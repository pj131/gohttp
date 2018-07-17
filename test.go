package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

type Order struct {
	Name string `json:"name"`
	Id   int
}

// 结构体转json
func structToJson() {
	var m = Order{"10010", 1}
	bytes, _ := json.Marshal(m)
	fmt.Printf("json %s\n", bytes)
}

type ss struct {
	S1 string
	S2 string
	S3 string
}

var s interface{} = &ss{
	S1: "s1",
	S2: "s2",
	S3: "s3",
}

func test() {
	val := reflect.ValueOf(s)
	fmt.Println(val)                      //&{s1 s2 s3}
	fmt.Println(val.Elem())               //{s1 s2 s3}
	fmt.Println(val.Elem().Field(0))      //s1
	val.Elem().Field(0).SetString("hehe") //S1大写
	fmt.Println(val.Elem().Interface())

	n := val.Elem().NumField()
	for i := 0; i < n; i++ {
		val.Elem().Field(i).SetString("hehe1") //S1大写
		fmt.Println(val.Elem().Interface())
	}
}

func dbtest() {
	db, err := sql.Open("mysql", "sa:cast1234@tcp(192.168.127.128:3306)/grid_new?charset=utf8")
	if err != nil {
		return
	}

	defer db.Close()

	rows, _ := db.Query("SELECT * FROM t_windplant")
	cols, _ := rows.Columns()
	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}

		//放入结果集
		result[i] = row
		i++
	}

	for item, resultmap := range result {
		fmt.Println(item, resultmap)
	}

}
