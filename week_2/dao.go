package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	sqlUser = "root"
	sqlPwd  = "1234567"
)

func InitSql() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", sqlUser+":"+sqlPwd+"@tcp(127.0.0.1:3306)/beego?charset=utf8")
	if err != nil {
		err = errors.Wrap(err, "sql Open fail")
		return
	}

	err = db.Ping()
	if err != nil {
		err = errors.Wrap(err, "sql ping fail")
		return
	}

	return
}

type user struct {
	Uid     int64
	Name    string
	Sex     int8
	Address string
}

func Query(db *sql.DB, sqlStr string) (err error) {
	fmt.Println("sqlStr:", sqlStr)
	rows, err := db.Query(sqlStr)
	if err != nil {
		return errors.Wrap(err, "db query fail")
	}

	for rows.Next() {
		var ac user
		err = rows.Scan(&ac.Uid, &ac.Name, &ac.Sex, &ac.Address)
		if err != nil {
			switch {
			case err == sql.ErrNoRows:
				err = errors.Wrap(err, "ErrNoRows")
			default:
				err = errors.Wrap(err, "scan error")
			}
			return
		}
		fmt.Println(ac)
	}

	fmt.Println("queryAc finished")

	return
}
