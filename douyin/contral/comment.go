package contral

import (
	"douyin/global"
	"douyin/model"
	"douyin/pkg"
	"douyin/service"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentREQ struct {
	User_id      int64  `json:"user_id" form:"user_id"`
	Token        string `json:"token" form:"token"`
	Video_id     int64  `json:"video_id" form:"video_id"`
	Action_type  int64  `json:"action_type" form:"action_type"`
	Comment_text string `json:"comment_text" form:"comment_text"`
	Comment_id   int64  `json:"comment_id" form:"comment_id"`
}

func CreateComment(c *gin.Context) {
	nowcommentid := time.Now().Unix()
	param := CommentREQ{}
	valid, errs := pkg.BindAndValid(c, &param)
	fmt.Println("the param:", param)
	if !valid {
		log.Println("CreatUserInDenglu bindandvalid err:", errs)
		c.JSON(500, gin.H{
			"status_code": -1,
		})
		return
	}
	if param.Action_type == 1 {
		pinglun := model.Comment{ID: int64(nowcommentid), Userid: param.User_id, Content: param.Comment_text, Videoid: param.Video_id}
		err := service.CreateComment(pinglun, global.DbEngine)
		nowcommentid++
		if err != nil {
			log.Println("create pinglun bindandvalid err:", errs)
			c.JSON(500, gin.H{
				"status_code": -1,
			})
			return
		}
		modelu := model.User{ID: param.User_id}
		user, err := service.GetUserbyId(modelu, global.DbEngine)
		if err != nil {
			log.Println("getuserby id  err:", err)
			log.Println("create pinglun bindandvalid err:", err)
			// c.JSON(500, gin.H{
			// 	"status_code": -1,
			// })
			// return
		}
		v := model.Video{ID: param.Video_id}
		video, err := service.GetVideo(v, global.DbEngine)
		if err != nil {
			log.Println("updatecomment getvideo   err:", err)
			log.Println("create pinglun bindandvalid err:", err)
			// c.JSON(500, gin.H{
			// 	"status_code": -1,
			// })
			// return
		}
		video.CommentCount += 1
		err = service.UpdateVideoComment(video, global.DbEngine)
		if err != nil {
			log.Println("updatecomment    err:", err)
			// c.JSON(500, gin.H{
			// 	"status_code": -1,
			// })
			// return
		}
		c.JSON(200, gin.H{
			"id":          param.Comment_id,
			"user":        user,
			"content":     param.Comment_text,
			"create_date": time.Now().Unix(),
		})

	}
}
func CommentList(c *gin.Context) {
	vid := c.Query("video_id")
	videoid, _ := strconv.Atoi(vid)
	modelc := model.Comment{Videoid: int64(videoid)}
	commentList, err := service.GetCommetList(modelc, global.DbEngine)
	if err != nil {
		log.Println("getcomment list  err:", err)
		c.JSON(500, gin.H{
			"status_code": -1,
		})
		return
	}
	c.JSON(200, gin.H{
		"status_code":  0,
		"status_msg":   "success",
		"comment_list": commentList,
	})
	return
}
