package service

import (
	"douyin/model"
	"log"

	"github.com/jinzhu/gorm"
)

func GetVideo(v model.Video, db *gorm.DB) (model.Video, error) {
	video, err := v.Get(db)
	if err != nil {
		log.Println(err)
		return video, err
	}
	return video, nil
}

func CreateVideo(v model.Video, db *gorm.DB) error {
	err := v.Create(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func GetVideoList(db *gorm.DB, lasttime int64) (*[]model.Videofeed, error, int64) {
	list, err, thetime := model.VideoList(db, lasttime)
	if err != nil {
		log.Println(err)
	}
	return list, nil, thetime
}
func GetdianzanList(db *gorm.DB) (*[]model.Videofeed, error) {
	list, err := model.Getdanzanlist(db)
	if err != nil {
		log.Println(err)
	}
	return list, nil
}

func UpdateVideoZan(v model.Video, db *gorm.DB) error {
	err := v.UpdateZan(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func UpdateVideoComment(v model.Video, db *gorm.DB) error {
	err := v.UpdateComment(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
