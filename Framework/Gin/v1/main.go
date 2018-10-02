package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

type Article struct {
	Id          uint   `gorm:"column:id; primary_key" json:"id"`
	GroupID     uint   `gorm:"column:g_id" json:"group_id"`
	TypeId      uint   `gorm:"column:t_id" json:"type_id"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description"  json:"description"`
	Content     string `gorm:"column:content"  json:"-"`

	ArticleGroup ArticleGroup `gorm:"ForeignKey:GroupID"`

	CreatedAt time.Time  `column:"created_at" json:"created_at"`
	UpdatedAt time.Time  `column:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `column:"deleted_at" json:"deleted_at"`
}

type ArticleGroup struct {
	Id   uint   `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name"`

	CreatedAt time.Time  `column:"created_at" json:"created_at"`
	UpdatedAt time.Time  `column:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `column:"deleted_at" json:"deleted_at"`
}

func (Article) TableName() string {
	return "blog_article"
}
func (ArticleGroup) TableName() string {
	return "blog_article_group"
}

func main() {
	db, err := gorm.Open("mysql", "demo:jia@tcp(120.79.134.74:22000)/blog?parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.LogMode(true)
	defer db.Close()

	r := gin.Default()

	r.GET("/article", func(c *gin.Context) {
		articles := []Article{}
		db.Find(&articles)
		c.JSON(http.StatusOK, gin.H{
			"data": articles,
		})
	})

	r.GET("/article/group", func(c *gin.Context) {
		ags := []ArticleGroup{}
		db.Find(&ags)
		for i := 0; i < len(ags); i++ {
			fmt.Println(ags[i].Name)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": ags,
		})
	})

	r.GET("/article/total", func(c *gin.Context) {
		//ags := []ArticleGroup{}
		article := Article{}
		db.Where("id = 1").Find(&article)
		db.Model(&article).Related(&article.ArticleGroup)

		//db.Table("blog_article_group").Select("blog_article_group.id, blog_article_group.name").Scan(&ags)

		//db.Preload("ArticleGroup").Find(&article)

		c.JSON(http.StatusOK, gin.H{
			"article": article,
		})
	})

	r.Run("127.0.0.1:9000") // listen and serve on 0.0.0.0:8080
}
