package main

import (
	"SimpleWebAPI/src/config"
	"SimpleWebAPI/src/handler"
	"SimpleWebAPI/src/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tkanos/gonfig"
)

var db *gorm.DB

func main() {

	// * declearing here to user later
	var errr error

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

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/", handler.PingGet())
		//v1.GET("/users", handler.UsersGet())
		v1.POST("/users", handler.UserPost(db))
		// v1.GET("/users/:id", fetchSingleTodo)
		// v1.PUT("/users/:id", updateTodo)
		// v1.DELETE("/users/:id", deleteTodo)
	}

	// * listen and serve on 0.0.0.0:8080 - default
	router.Run()
}
