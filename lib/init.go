package ms2mysql

import (
	"log"
	"os"
	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	// import odbc driver
	_ "github.com/lunny/godbc"
)

// MsEngine mssql引擎
var MsEngine *xorm.Engine
// MyEngine mysql引擎
var MyEngine *xorm.Engine
// Logger 日志公共变量
var Logger *log.Logger

func init() {
	f, err := os.OpenFile("sql.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		log.Fatalln("fail to create sql.log file!")
	}
	Logger = log.New(f, "", log.LstdFlags|log.Llongfile)
	//192.168.10.172
	MsEngine, err = xorm.NewEngine("odbc", "driver={SQL Server};Server=127.0.0.1;Database=kaoqin;uid=sa;pwd=123456;")
	if err != nil {
		log.Println("新建mssql引擎", err)
		return
	}
	if err = MsEngine.Ping(); err != nil {
		log.Println(err)
	}
	MsEngine.ShowSQL(true)
	MsEngine.Logger().SetLevel(core.LOG_DEBUG)
	//
	MyEngine, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/intro_oa?charset=utf8")
	if err != nil {
		log.Println("新建mysql引擎", err)
		return
	}
	if err = MyEngine.Ping(); err != nil {
		log.Println(err)
	}
	MyEngine.ShowSQL(true)
	MyEngine.Logger().SetLevel(core.LOG_DEBUG)
}
