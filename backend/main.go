package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Danmu struct {
	gorm.Model `json:"-"`
	Rgb        string `json:"rgb"`
	Content    string `json:"content"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("danmu.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Danmu{})

	//这里往下是更改弹幕内容部分
	for i := 1; i <= 15; i++ {
		db.Create(&Danmu{Rgb: "#39C5BB", Content: "我超 初音未来"})
	}
	for i := 15; i <= 30; i++ {
		db.Create(&Danmu{Rgb: "#66CCFF", Content: "我超 洛天依"})
	}
	//这里往上是更改弹幕内容部分

	r := gin.Default()

	r.GET("/danmu", func(c *gin.Context) {
		var danmus []Danmu
		db.Find(&danmus)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(200, danmus)
		//for i := 1; i <= 30; i++ {
		//	c.JSON(200, danmus[i])
		//}
	})

	r.Run(":8080")
}
