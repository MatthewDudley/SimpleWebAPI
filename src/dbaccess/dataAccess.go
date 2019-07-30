package dbaccess

import (
	"SimpleWebAPI/src/config"
	"SimpleWebAPI/src/model"
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// need in da file
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const fileName string = "dataAccess"

// InitializeDB makes the db connection | params: none | return: *sql.DB
func InitializeDB() *sql.DB {
	// * complete

	// * variables
	const configFilePath string = "./src/config/config.dev.json"
	var err error
	var db *sql.DB
	var conf config.Configuration // * allocate a DBConfig struct called configuration

	// * open and read config file into the conf struct
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalf("%s - ERROR occured at os.Open(cofigFilePath) line 29: %s", fileName, err)
	}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatalf("%s - ERROR occured at decoder.Decode(...) line 34: %s", fileName, err)
	}
	log.Printf("%s - Config file successfully read...", fileName)

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
		log.Fatalf("%s - ERROR occured at os.Open('mysq', connectionString) line 52: %s", fileName, err)
	}
	// * ping db to see if you have an established connection
	err = db.Ping()
	if err != nil {
		log.Printf("%s - ERROR db.Ping() line 72: %s", fileName, err)
	}
	log.Printf("%s - db.Ping() successful, established connection to db...", fileName)
	return db
}

// InsertUser params: model.User, *sql.DB | return: id int
func InsertUser(user *model.User, db *sql.DB) int64 {
	// * save the user to the db
	stmt := `INSERT INTO users (name, age) VALUES (?, ?)`
	res, err := db.Exec(stmt, user.Name, user.Age)
	if err != nil {
		log.Fatalf("%s - ERROR occured at db.Exec(...) line 69 %s", fileName, err)
	}
	// * get the last id inserted and store in id to return
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("%s - ERROR occured at res.LastInsertID() line 74 %s", fileName, err)
	}
	return id
}

// GetUsers params: user array and db ref
func GetUsers(user []model.User, db *sql.DB) error {
	// TODO complete db work to get all users
	// * select all users and return back in users array

	return nil
}
