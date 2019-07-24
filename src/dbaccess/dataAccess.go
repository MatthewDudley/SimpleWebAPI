package dbaccess

import (
	"SimpleWebAPI/src/config"
	"SimpleWebAPI/src/model"
	"fmt"

	"github.com/jinzhu/gorm"
	// need in da file
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tkanos/gonfig"
)

// InitializeDB makes the db connection
func InitializeDB() *gorm.DB {
	// * variables used
	var errr error
	var db *gorm.DB

	// * allocate a DBConfig struct called configuration
	var configuration config.DBConfiguration

	// * either read config files manually or use Gonfig below
	/* file, err := os.Open("./api/config/config.dev.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	} */

	// * read config files using gonfig and store them in the configuration struct
	errr = gonfig.GetConf("./src/config/config.dev.json", &configuration)
	if errr != nil {
		panic(errr)
	} else {
		fmt.Println("Successfully read configuration file...")
	}

	/*
				config.dev.json

		PATH: api/config/config.dev.json

		CONTENTS:
		{
			"dbname": "[DB NAME]",
			"dbuser": "[USERNAME]",
			"dbpassword": "[PASSWORD]"
		}

	*/
	// * build connection string from config.dev.json that was put into configuration struct above
	// * config.dev.json format is above
	connectionString := configuration.DBUser + ":" + configuration.DBPassword + "@/" + configuration.DBName

	// * open db connection using gorm
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to db...")
	}

	//Migrate the schema
	db.AutoMigrate(&model.User{})

	return db
}

// InsertUser params: user model.User
func InsertUser(user *model.User, db *gorm.DB) {
	// * save the user to the db using gorm
	db.Save(&user)
}
