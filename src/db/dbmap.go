package db

import (
	"database/sql"
	. "entity"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
)

var Dbmap *gorp.DbMap

func init() {
	Dbmap = initDb()
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", "test:test@/test")
	checkErr(err, "mysql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}
