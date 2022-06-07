package contral

import (
	"douyin/global"
	"douyin/model"
	"douyin/pkg"
	"douyin/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Zanrequest struct {
	User_id     string `json:"user_id" form:"user_id"`
	Token       string `json:"token"  form:"token"`
	Video_id    int64  `json:"video_id"  form:"video_id"`
	Action_type int64  `json:"action_type" form:"action_type"`
}

func UpdateZan(c *gin.Context) {
	fmt.Println("test updatezan")
	param := Zanrequest{}
	valid, errs := pkg.BindAndValid(c, &param)
	fmt.Println(param)
	if !valid {
		log.Println("CreatUserInDenglu bindandvalid err:", errs)
		c.JSON(200, gin.H{
			"status_code": -1,
		})
		return
	}
	var dianzan bool
	var zannum int64
	if param.Action_type == 1 {
		dianzan = true
		zannum = -1
	} else {
		dianzan = false
		zannum = 1
	}
	video, err := service.GetVideo(model.Video{ID: param.Video_id}, global.DbEngine)
	video.FavoriteCount += zannum
	video.IsFavorite = dianzan
	if err != nil {
		log.Println("update zan get video err:", err)
		c.JSON(200, gin.H{
			"status_code": -1,
		})
		return
	}
	err = service.UpdateVideoZan(video, global.DbEngine)
	if err != nil {
		log.Println("update zan update video err:", err)
		c.JSON(200, gin.H{
			"status_code": -1,
		})
		return
	}
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "sucess",
	})
	return

}
