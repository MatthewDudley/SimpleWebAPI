package handler

import (
	"SimpleWebAPI/src/dbaccess"
	"SimpleWebAPI/src/model"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UsersGet function
func UsersGet(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// * array of users from the model
		var users []model.User

		// * get users with find command
		db.Find(&users)
		if len(users) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "No user found!"})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"data":   users})
		}
	}
}

// UserPost - post a User with a given 'name' - string and 'age' - int to the db
func UserPost(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		// * convert age from a string to int
		age, _ := strconv.Atoi(c.PostForm("age"))
		// * set the new user
		user := model.User{
			Name: c.PostForm("name"),
			Age:  age}
		// * call da with the user model to insert
		dbaccess.InsertUser(c, &user, db)

		// ! if return from dbaccess.InsertUser == nil
		// * return to user it was successfull
		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "User successfully added!",
			"userID":  user.ID})
	}
}
