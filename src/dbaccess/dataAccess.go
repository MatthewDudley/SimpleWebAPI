package dbaccess

import (
	"SimpleWebAPI/src/config"
	"SimpleWebAPI/src/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	// need in da file
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitializeDB makes the db connection | params: none | return: *sql.DB
func InitializeDB() *sql.DB {
	// * complete

	// * variables
	const fileName string = "dataAccess.go"
	const configFilePath string = "./src/config/config.dev.json"
	var err error
	var db *sql.DB
	var conf config.Configuration // * allocate a DBConfig struct called configuration

	// * open and read config file into the conf struct
	configFile, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s - Config file successfully read..", fileName)

	// * could read config files using gonfig
	// errr = gonfig.GetConf(config.ConfigPath, &config)
	// if errr != nil {
	// 	panic(errr)
	// } else {
	// 	fmt.Println("Successfully read configuration file...")
	// }

	// * build connection string from config.dev.json that was put into configuration struct above
	// * config.dev.json format is above
	connectionString := conf.DBUser + ":" + conf.DBPassword + "@/" + conf.DBName

	// * open db connection
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	// * ping db to see if you have an established connection
	err = db.Ping()
	if err != nil {
		fmt.Printf("%s - Ping failed... Could not establish connection - line 72", fileName)
	}
	fmt.Printf("%s - Ping Successfull... established connection to db", fileName)
	return db
}

// InsertUser params: model.User, *sql.DB | return: error
func InsertUser(c *gin.Context, user *model.User, db *sql.DB) error {
	// TODO complete db work to save the passed in user struct

	// * save the user to the db

	return nil
}

// GetUsers params: user array and db ref
func GetUsers(user []model.User, db *sql.DB) error {
	// TODO complete db work to get all users
	// * select all users and return back in users array

	return nil
}
