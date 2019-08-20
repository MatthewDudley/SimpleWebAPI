package handler

import (
	"SimpleWebAPI/src/dbaccess"
	"SimpleWebAPI/src/model"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const fileName string = "user_handler"

// UsersGet function to get all records in Users table
func UsersGet(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		// * define users array
		var users []model.User

		// * get users with find command
		users = dbaccess.GetUsers(db)
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
		rtn := dbaccess.InsertUser(&user, db)
		// * return to user it was successfull
		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "User successfully added!",
			"userID":  rtn})
	}
}

// UserGet function to get a single user record by id
func UserGet(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		// * define user array
		var user *model.User

		// * get id from header
		id, _ := strconv.Atoi(c.Param("id"))

		// * pass id to da function to get the user
		user = dbaccess.GetUserByID(id, db)

		// * check user name if empty
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "No user found!"})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"data":   &user})
		}
	}

}

// UserPut function to update a single user record by id
func UserPut(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		// * define user array
		var user *model.User

		// * get id from header
		id, _ := strconv.Atoi(c.Param("id"))

		// * pass id to da function to get the user
		user = dbaccess.GetUserByID(id, db)

		// * check user name if empty
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "No user found!"})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"data":   &user})
		}
	}

}
