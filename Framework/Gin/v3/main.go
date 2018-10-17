package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Article struct {
	ID      int    `gorm:"primary_key" json:"id"`
	GroupId int    `gorm:"column:g_id" json:"g_id"`
	Title   string `json:"title"`

	Group Group `gorm:"ForeignKey:GroupId" json:"group"`
}

type Group struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (Article) TableName() string {
	return "blog_article"
}
func (Group) TableName() string {
	return "blog_article_group"
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/demo?parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.LogMode(true)
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		article := []Article{}
		//db.Find(&article)
		db.Preload("Group").Find(&article)
		//for i := 0; i < len(article); i++ {
		//	db.Model(&article[i]).Related(&article[i].Group)
		//}

		c.JSON(http.StatusOK, gin.H{
			"data": article,
		})
	})

	r.Run("127.0.0.1:8000") // listen and serve on 0.0.0.0:8080
}
