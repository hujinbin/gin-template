package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDatabase() {
	var err error
	Db, err = sql.Open("mysql", Cfg.MysqlUser+":"+Cfg.MysqlPwd+"@tcp("+Cfg.MysqlHost+":"+Cfg.MysqlPort+")/"+Cfg.MysqlDb)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}