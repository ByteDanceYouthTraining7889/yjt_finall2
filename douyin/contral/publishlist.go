package contral

import (
	"douyin/global"
	"douyin/service"
	"log"

	"github.com/gin-gonic/gin"
)

func PublishList(c *gin.Context) {
	tags, err := service.GetPublishList(global.DbEngine)
	if err != nil {
		log.Println("viode list  err：", err)
		return
	}
	log.Println(tags)
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "sucess",
		"video_list":  tags, //*************
	})
	return
}
