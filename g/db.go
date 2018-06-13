package g

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	// _ "github.com/go-sql-driver/mysql"
	"srv-db/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/consul"
)

var DB *gorm.DB

type DatabaseConfig struct {
	Host     string `json:"dbhost"`
	User     string `json:"dbuser"`
	Password string `json:"dbpassword"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

func readDBConnString(dbconfig *DatabaseConfig, conn *string) error {
	var buffer bytes.Buffer
	if dbconfig.User == "" || dbconfig.Port == 0 || dbconfig.DBName == "" || dbconfig.Host == "" || dbconfig.Password == "" {
		return fmt.Errorf("read db config: dbconfig is nil")
	}
	buffer.WriteString(dbconfig.User)
	buffer.WriteString(":")
	buffer.WriteString(dbconfig.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(dbconfig.Host)
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(dbconfig.Port))
	buffer.WriteString(")/")
	buffer.WriteString(dbconfig.DBName)
	buffer.WriteString("?loc=Local&parseTime=true&charset=utf8")
	*conn = buffer.String()
	return nil
}

func InitDBConfig() {
	var err error
	var dbconfig DatabaseConfig
	var connectString string

	consulSource := consul.NewSource(
		// optionally specify consul address; default to localhost:8500
		consul.WithAddress(consulAdress),
		// optionally specify prefix; defaults to /micro/config
	)
	conf := config.NewConfig()

	// Load file source
	conf.Load(consulSource)
	conf.Get("micro", "config", "database").Scan(&dbconfig)
	// log.Println(dbconfig)
	readDBConnString(&dbconfig, &connectString)
	log.Println(connectString)
	// DB, err = sql.Open("mysql", connectString)
	DB, err := gorm.Open("mysql", connectString)
	// log.Println(DB, err)
	if err != nil {
		log.Fatal("open db fail", err)
	}
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetMaxIdleConns(10)
	log.Println("Initialized Db Configuration.")
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Users{}, &model.AliPay{}, &model.CreditCard{}, &model.Email{}, &model.HomeAddress{}, &model.ShoppingAdress{}, &model.Wechat{}, &model.Product{})
	// DB.Create(&model.Product{Code: "L1212", Price: 1000})
	// adresses := make([]model.HomeAddress, 0)
	// homeAddress := model.HomeAddress{Address: "Shanghai", Post: "djkjflsdjflsdjlf"}
	// adresses = append(adresses, homeAddress)

	// DB.Create(&model.Users{Name: "jamesduan", Sex: "男", Number: "12345667", HomeAddresses: adresses, BirthDay: time.Unix(1469579899, 0)})
}
