package model

import (
	"douyin/global"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Video struct {
	Authorid      int64  `json:"authorid"`       // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
	CreatedOn     int64  `json:"created_on"`
}
type Videofeed struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
	CreatedOn     int64  `json:"created_on" `
}

func NewVideo() Video {
	return Video{}
}

func (v Video) TableName() string {
	return "video"
}
func (v Video) Get(db *gorm.DB) (Video, error) {
	db = db.Where(
		"id=?",
		v.ID,
	)
	err := db.First(&v).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("video get err:", err)
		return v, err
	}
	return v, nil

}
func Getdanzanlist(db *gorm.DB) (*[]Videofeed, error) {
	var videolist []Video
	var err error

	db = db.Limit(30) //==============后面前

	db = db.AutoMigrate(&Video{}) //自己加的 自动绑定结构体
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	db = db.Where("is_favorite=?", true)
	if err = db.Debug().Find(&videolist).Error; err != nil {
		log.Println("find dianzanlist err:", err)
		return nil, err
	}
	var Feedlist []Videofeed
	for k, v := range videolist {
		u := User{ID: v.Authorid}
		fmt.Println("the k:", k)
		// Feedlist[k] = Videofeed{Author: u, CommentCount: v.CommentCount, FavoriteCount: v.FavoriteCount, ID: v.ID, IsFavorite: v.IsFavorite, CoverURL: v.CoverURL, PlayURL: v.PlayURL, Title: v.Title}
		Feedlist = append(Feedlist, Videofeed{Author: u, CommentCount: v.CommentCount, FavoriteCount: v.FavoriteCount, ID: v.ID, IsFavorite: v.IsFavorite, CoverURL: v.CoverURL, PlayURL: v.PlayURL, Title: v.Title})
	}
	return &Feedlist, nil
}
func VideoList(db *gorm.DB, lasttime int64) (*[]Videofeed, error, int64) {
	var videolist []Video
	var err error

	db = db.Limit(30) //==============后面前

	db = db.AutoMigrate(&Video{}) //自己加的 自动绑定结构体
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	// db = db.Where("state=?", t.)
	db = db.Where("created_on < ?", lasttime).Order("created_on desc")
	if err = db.Debug().Find(&videolist).Error; err != nil {
		log.Println("find videolist err:", err)
		return nil, err, time.Now().Unix()
	}
	l := len(videolist)
	var thetime int64
	if l != 0 {
		// thetime = videolist[l-1].CreatedOn  刷新评论才能出来 所以先把这个去掉 后面多了在看
		thetime = time.Now().Unix()
	} else {
		thetime = time.Now().Unix()
	}

	var Feedlist []Videofeed
	for _, v := range videolist {
		u := User{ID: v.Authorid}
		user, err := u.Getbyid(global.DbEngine)
		fmt.Println("is 1 =true ,", user.Is_follow)
		if err != nil {
			log.Println("find videolist err:", err)
		}
		// Feedlist[k] = Videofeed{Author: user, CommentCount: v.CommentCount, FavoriteCount: v.FavoriteCount, ID: v.ID, IsFavorite: v.IsFavorite, CoverURL: v.CoverURL, PlayURL: v.PlayURL, Title: v.Title}
		Feedlist = append(Feedlist, Videofeed{Author: user, CommentCount: v.CommentCount, FavoriteCount: v.FavoriteCount, ID: v.ID, IsFavorite: v.IsFavorite, CoverURL: v.CoverURL, PlayURL: v.PlayURL, Title: v.Title})
	}
	return &Feedlist, nil, thetime
}
func PublishList(db *gorm.DB) (*[]Videofeed, error) {
	var publishlist []Video
	var err error

	db = db.Limit(30) //==============后面前

	db = db.AutoMigrate(&Video{}) //自己加的 自动绑定结构体
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	// db = db.Where("state=?", t.)
	if err = db.Find(&publishlist).Error; err != nil {
		log.Println("find publishlist err:", err)
		return nil, err
	}
	var Feedlist []Videofeed
	for _, v := range publishlist {
		u := User{ID: v.Authorid}
		user, err := u.Getbyid(global.DbEngine)
		fmt.Println("the err:", err)
		// Feedlist[k] = Videofeed{Author: user, CommentCount: v.CommentCount, FavoriteCount: v.FavoriteCount, ID: v.ID, IsFavorite: v.IsFavorite, CoverURL: v.CoverURL, PlayURL: v.PlayURL, Title: v.Title}
		Feedlist = append(Feedlist, Videofeed{Author: user, CommentCount: v.CommentCount, FavoriteCount: v.FavoriteCount, ID: v.ID, IsFavorite: v.IsFavorite, CoverURL: v.CoverURL, PlayURL: v.PlayURL, Title: v.Title})
	}
	return &Feedlist, nil
}
func (v Video) Create(db *gorm.DB) error {
	return db.Debug().Create(&v).Error
}

func (v Video) UpdateZan(db *gorm.DB) error {
	db = db.Model(&User{}).Where("id=?", v.ID) //model 指定db操作的实力模型 一般在没有find 等指定类型时使用 操作限制条件是传进来的tag的数据行
	//fmt.Println(t.Name)
	err := db.Select("is_favorite").Save(v).Error //用save不报错 create报错 只能先将就save 用了 配合selct 也解决了 0值不能改的问题
	if err != nil {
		fmt.Println("video zan 更新错误：", err)
		return err
	}
	return nil
}
func (v Video) UpdateComment(db *gorm.DB) error {
	db = db.Model(&User{}).Where("id=?", v.ID) //model 指定db操作的实力模型 一般在没有find 等指定类型时使用 操作限制条件是传进来的tag的数据行
	//fmt.Println(t.Name)
	err := db.Debug().Select("comment_count").Save(v).Error //用save不报错 create报错 只能先将就save 用了 配合selct 也解决了 0值不能改的问题
	if err != nil {
		fmt.Println("video comment 更新错误：", err)
		return err
	}
	return nil
}
