package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	Username      string `gorm:"primary key" json:"username" form:"username" `
	Password      string `json:"password" form:"password"`
	Is_follow     bool   `json:"is_follow"`
}

func NewUser() User {
	return User{}
}

func (u User) TableName() string {
	return "user"
}
func (u User) Get(db *gorm.DB) (User, error) {
	fmt.Println("test user:", u)
	db = db.Debug().Where( //debug 会在终端显示
		"username=? AND password=? ",
		u.Username, u.Password,
	)
	err := db.First(&u).Error
	if err != nil {
		fmt.Println("user get err:", err)
		return u, err
	}
	return u, nil

}
func (u User) Getbyid(db *gorm.DB) (User, error) {
	fmt.Println("test user:", u)
	db = db.Debug().Where( //debug 会在终端显示
		"id=?",
		u.ID,
	)
	err := db.First(&u).Error
	if err != nil {
		fmt.Println("user getbyid  err:", err)
		return u, err
	}
	fmt.Println("test1 user:", u)
	return u, nil
}
func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}
func (u User) Update(db *gorm.DB) error {
	db = db.Model(&User{}).Where("username=?", u.Username) //model 指定db操作的实力模型 一般在没有find 等指定类型时使用 操作限制条件是传进来的tag的数据行
	//fmt.Println(t.Name)
	err := db.Select("*").Save(&u).Error //用save不报错 create报错 只能先将就save 用了 配合selct 也解决了 0值不能改的问题
	if err != nil {
		fmt.Println("user 更新错误：", err)
		return err
	}
	return nil
}
func (u User) UpdateGuanzhu(db *gorm.DB) error {
	db = db.Model(&User{}).Where("id=?", u.ID) //model 指定db操作的实力模型 一般在没有find 等指定类型时使用 操作限制条件是传进来的tag的数据行
	fmt.Println("这是")
	err := db.Debug().Select("follow_count", "is_follow").Update(&u).Error //用save不报错 create报错 只能先将就save 用了 配合selct 也解决了 0值不能改的问题**********************
	if err != nil {
		fmt.Println("user 更新错误：", err)
		return err
	}
	return nil
}
