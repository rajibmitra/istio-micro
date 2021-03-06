package cinit

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

var Mysql *sqlx.DB

//初始化连接
func mysqlInit() {
	var err error
	dataSourceName := Config.Mysql.User + ":" + Config.Mysql.Password + "@tcp(" + Config.Mysql.Addr + ":" + strconv.Itoa(Config.Mysql.Port) +
		")/" + Config.Mysql.DbName + "?parseTime=true&loc=Local"
	Mysql, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	Mysql.SetMaxIdleConns(Config.Mysql.IdleConn)
	Mysql.SetMaxOpenConns(Config.Mysql.MaxConn)
	err = Mysql.Ping()
	if err != nil {
		panic(err)
	}
}

//关闭
func mysqlClose() {
	Mysql.Close()
}
