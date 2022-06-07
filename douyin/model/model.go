package model

import (
	"douyin/setting"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //一定不要忘记*********************
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) { //数据库创建链接 用了配置方法
	//s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local")
	var err error
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", databaseSetting.UserName, databaseSetting.Password, databaseSetting.Host, databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)) //有关数据库的信息来连接
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	return db, nil
}
