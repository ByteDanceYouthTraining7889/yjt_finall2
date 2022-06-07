package model

import (
	"douyin/global"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// Comment
type Comment struct {
	Content    string `json:"content" form:"content"`         // 评论内容
	CreateDate string `json:"create_date" form:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64  `json:"id" form:"id"`                   // 评论id
	Userid     int64  `json:"userid" form:"userid"`           // 评论用户信息
	Videoid    int64  `json:"videoid" form:"videoid"`
}
type CommentResp struct {
	Content    string `json:"content" form:"content"`          // 评论内容
	CreateDate string `json:"create_date"  form:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64  `gorm:"primary key" json:"id" form:"id"` // 评论id
	User       User   `json:"user" form:"user"`                // 评论用户信息
}

func (c Comment) TableName() string {
	return "comment"
}
func (c Comment) GetComment(db *gorm.DB) (Comment, error) {
	db = db.Debug().Where( //debug 会在终端显示
		"id=?",
		c.ID,
	)
	err := db.First(&c).Error
	if err != nil {
		fmt.Println("comment get err:", err)
		return c, err
	}
	return c, nil
}
func (c Comment) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}
func (c Comment) Delete(db *gorm.DB) error {
	return db.Debug().Where("id=?", c.ID).Delete(&c).Error
}
func CommentList(c Comment, db *gorm.DB) (*[30]CommentResp, error) {
	var Commentlist []Comment
	var err error

	db = db.Limit(30) //==============后面前

	db = db.AutoMigrate(&Comment{}) //自己加的 自动绑定结构体
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	db = db.Where("videoid=?", c.Videoid)
	if err = db.Find(&Commentlist).Error; err != nil {
		log.Println("find commentlist err:", err)
		return nil, err
	}
	var Comlist [30]CommentResp
	for k, v := range Commentlist {
		u := User{ID: v.Userid}
		user, err := u.Getbyid(global.DbEngine)
		log.Println("find videolist err:", err)
		Comlist[k] = CommentResp{Content: v.Content, CreateDate: v.CreateDate, ID: v.ID, User: user}
	}
	return &Comlist, nil
}
