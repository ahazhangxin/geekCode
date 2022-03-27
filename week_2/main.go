package main

import (
	"fmt"
)

func main() {
	sql, err := InitSql()
	defer sql.Close()
	if err != nil {
		fmt.Println("initSql error:", err)
		fmt.Printf("%+v\n", err)
	}

	err = Query(sql, "select * from user where id=10001")
	if err != nil {
		fmt.Println("query sql error :", err)
		fmt.Printf("%+v\n", err)
	}
}
