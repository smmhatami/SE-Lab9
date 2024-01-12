package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init_db() {
	fmt.Println("initializing database")
	dsn := "host=postgres_db user=postgres password=postgres port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	// Create Users
	db.Create(&User{ID: 1, FirstName: "ali", LastName: "alavi"})

}

func main() {
	dsn := "host=postgres_db user=postgres password=postgres port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var users []User
	db.Raw("SELECT * FROM users").Scan(&users)
	if len(users) == 0 {
		init_db()
	}

	r := gin.Default()

	r.GET("/user/", GetUsers)
	r.GET("/user/:id", GetUser)
	r.POST("/user", CreateUser)
	r.PUT("/user/:id", UpdateUser)
	r.DELETE("/user/:id", DeleteUser)

	r.Run(":8080")
}

func GetUsers(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	d := db.Where("id = ?", id).Delete(&user)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
