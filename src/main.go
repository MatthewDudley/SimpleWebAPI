package main

import (
	"SimpleWebAPI/src/dbaccess"
	"SimpleWebAPI/src/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {

	// make a call to connect to db
	db = dbaccess.InitializeDB()

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
