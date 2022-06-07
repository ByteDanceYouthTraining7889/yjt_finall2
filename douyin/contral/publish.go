package contral

import (
	"douyin/global"
	"douyin/model"
	"douyin/pkg"
	"douyin/service"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func FabuVideo(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("data")
	token := c.PostForm("token")
	if token == "" {
		log.Println("publish token get err:", err)
		return
	}
	claims, err := pkg.ParseToken(token)
	if err != nil {
		log.Println("parse token  err:", err)
	}
	u := model.User{Username: claims.Id, Password: claims.Secret, FollowCount: 0, FollowerCount: 0}
	user, err := service.GetUser(u, global.DbEngine)
	if err != nil {
		log.Println("fabuvedio user get   err:", err)
	}
	fileType := 10
	if err != nil {
		log.Println("filetype err:", err)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		log.Println("fileHeader err:", err)
		return
	}
	fileInfo, err := service.UploadFile(pkg.FlieType(fileType), file, fileHeader)
	if err != nil {
		log.Println("file uploadfile err:", err)
		return

	}
	Videoid := time.Now().Unix()
	nowvideo := model.Video{Authorid: user.ID, PlayURL: fileInfo.AccessUrl, ID: int64(Videoid), CreatedOn: time.Now().Unix()}

	err = service.CreateVideo(nowvideo, global.DbEngine)
	if err != nil {
		log.Println("video create err:", err)
		return
	}
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"url":         fileInfo.AccessUrl,
	})

}
