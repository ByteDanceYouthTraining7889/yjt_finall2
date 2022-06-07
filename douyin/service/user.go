package service

import (
	"douyin/model"
	"log"

	"github.com/jinzhu/gorm"
)

func GetUser(u model.User, db *gorm.DB) (model.User, error) {
	user, err := u.Get(db)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}
func GetUserbyId(u model.User, db *gorm.DB) (model.User, error) {
	user, err := u.Getbyid(db)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func CreateUser(u model.User, db *gorm.DB) error {
	err := u.Create(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func UpdateUser(u model.User, db *gorm.DB) error {
	err := u.Update(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func UpdateUserguanzhu(u model.User, db *gorm.DB) error {
	err := u.UpdateGuanzhu(db)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func GetuserGuanzhuList(g model.Guanzhu, db *gorm.DB) (*[]model.User, error) {
	list, err := g.UserguanzhuList(db)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return list, nil
}
func GetuserFensiList(g model.Guanzhu, db *gorm.DB) (*[]model.User, error) {
	list, err := g.UserFensiList(db)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return list, nil
}
