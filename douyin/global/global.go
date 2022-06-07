package global

import (
	"douyin/setting"

	"github.com/jinzhu/gorm"
)

var (
	DbEngine      *gorm.DB
	JWTSettings   setting.JWTSetting
	IndexPageSize setting.PinglunindexPager
	XiangQingSize setting.PinglunXiangqingPager
	FileSave      setting.UploadFileSetting
)
