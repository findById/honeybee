package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	addr     = "127.0.0.1:3306" //IP地址
	username = "root"           //用户名
	password = "root"           //密码
	dbName   = "iot"            //表名
)

var db *sql.DB

func init() {
	source := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=UTF8", username, password, addr, dbName)
	var err error = nil
	db, err = sql.Open("mysql", source)
	if err != nil {
		log.Println(err)
	}
}

func GetDBM() *sql.DB {
	return db
}
