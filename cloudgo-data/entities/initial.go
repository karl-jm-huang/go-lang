package entities

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

// use xorm instead of mysql
func init() {
	Engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	err = Engine.Sync2(new(UserInfo))
	checkErr(err)
	engine = Engine
}

// var mydb *sql.DB

// func init() {
// 	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
// 	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
// 	if err != nil {
// 		panic(err)
// 	}
// 	mydb = db
// }

// SQLExecer interface for supporting sql.DB and sql.Tx to do sql statement
// type SQLExecer interface {
// 	Exec(query string, args ...interface{}) (sql.Result, error)
// 	Prepare(query string) (*sql.Stmt, error)
// 	Query(query string, args ...interface{}) (*sql.Rows, error)
// 	QueryRow(query string, args ...interface{}) *sql.Row
// }

// DaoSource Data Access Object Source
// type DaoSource struct {
// 	// if DB, each statement execute sql with random conn.
// 	// if Tx, all statements use the same conn as the Tx's connection
// 	SQLExecer
// }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
