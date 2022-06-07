package contral

import (
	"douyin/global"
	"douyin/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetVideoFeed(c *gin.Context) {
	latest_time := c.Query("latest_time")
	thetime, err := strconv.Atoi(latest_time)
	tags, err, next_time := service.GetVideoList(global.DbEngine, int64(thetime))
	if err != nil {
		log.Println("viode list  err：", err)
		return
	}
	log.Println(tags)
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "sucess",
		"next_time":   next_time, //后面改进
		"video_list":  tags,      //*************
	})
	return
}
