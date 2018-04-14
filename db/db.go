package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDBConfig() {
	var err error
	DB, err = sql.Open("mysql", "root:jd123@tcp(127.0.0.1:3306)/weiping?loc=Local&parseTime=true")
	if err != nil {
		log.Fatal("open db fail", err)
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(15)
	log.Println("Initialized Db Configuration.")
}
