package handler

import (
	"SimpleWebAPI/src/dbaccess"
	"SimpleWebAPI/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UsersGet function
func UsersGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"GET": "Users",
		})
	}
}

// UserPost - post a User with a given 'name' - string and 'age' - int to the db
func UserPost(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		// * convert age from a string to int
		age, _ := strconv.Atoi(c.PostForm("age"))
		// * set the new user
		user := model.User{Name: c.PostForm("name"), Age: age}
		// * call da with the user model to insert
		dbaccess.InsertUser(&user, db)
		// * return to user it was successfull
		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "User successfully added!",
			"userID":  user.ID})
	}
}
