package service

import (
	"douyin/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func Getguanzhu(g model.Guanzhu, db *gorm.DB) (model.Guanzhu, error) {
	guanzhu, err := g.Get(db)
	if err != nil {
		fmt.Println("get guanzhu err:", err)
		return guanzhu, err
	}
	return guanzhu, nil
}
func Createguanzhu(g model.Guanzhu, db *gorm.DB) error {
	err := g.Create(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func UpdateGuanzhu(g model.Guanzhu, db *gorm.DB) error {
	err := g.UpdateGuanzhu(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func DeleteGuanzhu(g model.Guanzhu, db *gorm.DB) error {
	err := g.DeleteGuanzhu(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
