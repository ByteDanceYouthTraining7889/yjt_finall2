package contral

import (
	"douyin/global"
	"douyin/pkg"
	"douyin/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type DianzanlistRequest struct {
	User_id int64  `json:"user_id" form:"user_id"`
	Token   string `json:"token" form:"token"`
}

func DianzanList(c *gin.Context) {
	param := DianzanlistRequest{}
	valid, errs := pkg.BindAndValid(c, &param)
	fmt.Println(param)
	if !valid {
		log.Println("dianzanrequest bindandvalid err:", errs)
		c.JSON(200, gin.H{
			"status_code": -1,
		})
		return
	}
	tags, err := service.GetdianzanList(global.DbEngine)
	if err != nil {
		log.Println("viode dianzanlist  err：", err)
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
