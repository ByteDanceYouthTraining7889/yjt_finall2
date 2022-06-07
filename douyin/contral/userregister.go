package contral

import (
	"douyin/global"
	"douyin/model"
	"douyin/pkg"
	"douyin/service"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatUserInDenglu(c *gin.Context) {
	nowid := time.Now().Unix()
	param := model.NewUser()
	valid, errs := pkg.BindAndValid(c, &param)
	fmt.Println(param)
	if !valid {
		log.Println("CreatUserInDenglu bindandvalid err:", errs)
		c.JSON(200, gin.H{
			"status_code": -1,
		})
		return
	}
	param.ID = int64(nowid)
	nowid += 1
	err := service.CreateUser(param, global.DbEngine)
	if err != nil {
		log.Println("CreatUserInDenglu CreateUser err:", errs)
		c.JSON(200, gin.H{
			"status_code": -1,
		})
		return
	}
	token, err := pkg.GenerateToken(param.Username, param.Password) //生成token
	if err != nil {
		log.Println(" GenerateToken err:", err)
		c.JSON(200, gin.H{
			"status_code": -1,
		})
	}
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "创建成功",
		"user_id":     param.ID,
		"token":       token,
	})
	return
}
