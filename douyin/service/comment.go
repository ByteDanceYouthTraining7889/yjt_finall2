package service

import (
	"douyin/model"
	"log"

	"github.com/jinzhu/gorm"
)

func GetComment(c model.Comment, db *gorm.DB) (model.Comment, error) {
	comment, err := c.GetComment(db)
	if err != nil {
		log.Println(err)
		return comment, err
	}
	return comment, nil
}
func GetCommetList(c model.Comment, db *gorm.DB) (*[30]model.CommentResp, error) {
	list, err := model.CommentList(c, db)
	if err != nil {
		log.Println(err)
	}
	return list, nil
}
func CreateComment(c model.Comment, db *gorm.DB) error {
	err := c.Create(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func DeleteComment(c model.Comment, db *gorm.DB) error {
	err := c.Delete(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
