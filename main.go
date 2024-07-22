package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/sunkuo/libs/tools"
)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// a := libss.UtilFunction()
	// log.Fatal("!!!!!!!---", a, "----!!!!!")

	b := tools.PrintStr()
	log.Fatal("!!!!!!!---", b, "----!!!!!")

	// 資料庫連接
	dsn := "user:password@tcp(db:3306)/myapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自動遷移（創建表格）
	db.AutoMigrate(&User{})

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		var users []User
		result := db.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")
}
