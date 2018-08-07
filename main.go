package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// User model
type User struct {
	gorm.Model
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Test": "success"})
	fmt.Println("Root endpoint hit")
}

func createUser(c *gin.Context) {
	var user User     // create a ref User struct in mem
	c.BindJSON(&user) // bind router context data to User ref as JSON
	db.Create(&user)  // create User table in db
	c.JSON(200, user) // send server response
	return
}

func deleteUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	db.Where("id = ?", id).Find(&user)
	db.Delete(user)
	return
}

func getAllUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
	return
}

func updateUser(c *gin.Context) {
	username := c.Query("username")
	email := c.Query("email")
	user := User{UserName: username, Email: email}
	db.Where("email = ?", email).Find(&user)
	user.UserName = username
	db.Save(user)
	return
}

func seedUsers() {
	u1 := User{UserName: "tim", Email: "t@t.com"}
	u2 := User{UserName: "tay", Email: "s@t.com"}
	u3 := User{UserName: "tom", Email: "w@t.com"}
	db.Create(&u1)
	db.Create(&u2)
	db.Create(&u3)
}

func main() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	//seedUsers()           // uncomment to add 3 users as seed data
	r := gin.Default()
	r.GET("/", homePage)
	r.GET("/users", getAllUsers)             // /users Return all users
	r.POST("/user/new", createUser)          // /user/new Create a new user
	r.DELETE("/user/delete/:id", deleteUser) // /user/delete/?name Delete a user
	r.PUT("/user/change/", updateUser)       // Update a users email
	r.Run()                                  // Listen and serve on :8080

}
