package model

import (
	"douyin/global"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Guanzhu struct {
	Userid     int64 `gorm:"primary key" json:"userid" form:"userid"`
	Touserid   int64 `gorm:"primary key" json:"to_user_id" form:"to_user_id"`
	Actiontype int64 `json:"action_type" form:"action_type"`
}

func (g Guanzhu) TableName() string {
	return "guanzhu"
}
func (g Guanzhu) Get(db *gorm.DB) (Guanzhu, error) {
	fmt.Println("test user:", g)
	db = db.Debug().Where( //debug 会在终端显示
		"userid=? and to_user_id=?", g.Userid, g.Touserid,
	)
	err := db.First(&g).Error
	if err != nil {
		fmt.Println("guanzhu get err:", err)
		return g, err
	}
	return g, nil

}

func (g Guanzhu) Create(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}
func (g Guanzhu) UpdateGuanzhu(db *gorm.DB) error {
	db = db.Model(&Guanzhu{}).Where("userid=? and to_user_id=?", g.Userid, g.Touserid) //model 指定db操作的实力模型 一般在没有find 等指定类型时使用 操作限制条件是传进来的tag的数据行
	//fmt.Println(t.Name)
	err := db.Select("action_type").Save(g).Error //用save不报错 create报错 只能先将就save 用了 配合selct 也解决了 0值不能改的问题
	if err != nil {
		log.Println("guanzhu 更新错误：", err)
		return err
	}
	return nil
}
func (g Guanzhu) DeleteGuanzhu(db *gorm.DB) error {
	return db.Debug().Delete(&g).Error
}
func (g Guanzhu) UserguanzhuList(db *gorm.DB) (*[]User, error) {
	var guanzhulist []Guanzhu
	var err error

	db = db.Limit(30) //==============后面前

	db = db.AutoMigrate(&Guanzhu{}) //自己加的 自动绑定结构体
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	db = db.Where("userid=?", g.Userid)
	if err = db.Find(&guanzhulist).Error; err != nil {
		log.Println("find userlist err:", err)
		return nil, err
	}
	var guanzhuuserlist []User
	for _, v := range guanzhulist {
		nowu := User{ID: v.Touserid}
		fmt.Println("nowu:", nowu)
		user, err := nowu.Getbyid(global.DbEngine)
		if err != nil {
			log.Println("find user when err:", err)
			return nil, err
		}
		// guanzhuuserlist[k] = user
		guanzhuuserlist = append(guanzhuuserlist, user)
	}
	return &guanzhuuserlist, nil
}
func (g Guanzhu) UserFensiList(db *gorm.DB) (*[]User, error) {
	var guanzhulist []Guanzhu
	var err error

	db = db.Limit(30) //==============后面前

	db = db.AutoMigrate(&Guanzhu{}) //自己加的 自动绑定结构体
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	db = db.Where("touserid=?", g.Touserid)
	if err = db.Debug().Find(&guanzhulist).Error; err != nil {
		log.Println("find userfensilist err:", err)
		return nil, err
	}
	log.Println("test fensi guanzhulist test :", guanzhulist)
	var guanzhuuserlist []User
	for _, v := range guanzhulist {
		nowu := User{ID: v.Userid}
		user, err := nowu.Getbyid(global.DbEngine)
		if err != nil {
			log.Println("find user when fensi err:", err)
		}
		// guanzhuuserlist[k] = user
		guanzhuuserlist = append(guanzhuuserlist, user)
	}
	return &guanzhuuserlist, nil
}
