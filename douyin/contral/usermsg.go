package contral

import (
	"douyin/global"
	"douyin/model"
	"douyin/pkg"
	"douyin/service"
	"log"

	"github.com/gin-gonic/gin"
)

type UserMsg struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

func GetUserMsg(c *gin.Context) {
	param := UserMsg{}
	//response := app.NewResponse(c)
	valid, errs := pkg.BindAndValid(c, &param) //获取与验证
	if !valid {
		log.Println("bindandvalid errs :%v", errs)
		c.JSON(500, gin.H{
			"status_code": -1,
		})
		return
	}
	if param.Token == "" {
		log.Println("No token errs :%v", errs)
		c.JSON(500, gin.H{
			"status_code": -1,
		})
		return
	}
	paratoken, err := pkg.ParseToken(param.Token)
	if err != nil {
		log.Println("parsetoken err:", err)

	}
	u := model.User{Username: paratoken.Id, Password: paratoken.Secret}
	log.Println("is parsetoken correct?", u)
	user, err := service.GetUser(u, global.DbEngine)
	if err != nil {
		log.Println("GEtUser err:", err)
		c.JSON(500, gin.H{
			"status_code": -1,
		})
		return
	}
	// fmt.Println("token is")
	// fmt.Println(token)
	c.JSON(200, gin.H{
		"user":        user,
		"status_code": 0,
		"status_msg":  "sucess",
	})

}
