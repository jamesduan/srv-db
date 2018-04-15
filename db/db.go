package db

import (
	"bytes"
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/consul"
)

var DB *sql.DB

type DatabaseConfig struct {
	Host     string `json:"dbhost"`
	User     string `json:"dbuser"`
	Password string `json:"dbpassword"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

func readDBConnString(dbconfig *DatabaseConfig) string {
	var buffer bytes.Buffer
	buffer.WriteString(dbconfig.User)
	buffer.WriteString(":")
	buffer.WriteString(dbconfig.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(dbconfig.Host)
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(dbconfig.Port))
	buffer.WriteString(")/")
	buffer.WriteString(dbconfig.DBName)
	buffer.WriteString("?loc=Local&parseTime=true")
	return buffer.String()
}

func InitDBConfig() {
	var err error
	var dbconfig DatabaseConfig

	consulSource := consul.NewSource(
		// optionally specify consul address; default to localhost:8500
		consul.WithAddress("127.0.0.1:8500"),
		// optionally specify prefix; defaults to /micro/config
	)
	conf := config.NewConfig()

	// Load file source
	conf.Load(consulSource)
	conf.Get("micro", "config", "database").Scan(&dbconfig)
	// log.Println(dbconfig)
	DB, err = sql.Open("mysql", readDBConnString(&dbconfig))
	if err != nil {
		log.Fatal("open db fail", err)
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(15)
	log.Println("Initialized Db Configuration.")
}
