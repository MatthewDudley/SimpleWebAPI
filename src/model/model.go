package model

// User ID int, Name string, Age int
type User struct {
	ID   int
	Name string
	Age  int
}

// * named User due to my table being named users, gorm uses plural names
// * is named UserModel gorm would have looked for user_models etc. for other names
