package contral

import (
	"douyin/global"
	"douyin/model"
	"douyin/pkg"
	"douyin/service"
	"log"

	"github.com/gin-gonic/gin"
)

func UserInlogin(c *gin.Context) {
	param := model.User{}
	//response := app.NewResponse(c)
	username := c.Query("username")
	password := c.Query("password")
	param.Username = username
	param.Password = password

	user, err := service.GetUser(param, global.DbEngine)
	if err != nil {
		log.Println("GEtUser err:", err)
		c.JSON(500, gin.H{
			"status_code": -1,
		})
		return
	}
	token, err := pkg.GenerateToken(user.Username, user.Password) //生成token
	if err != nil {
		log.Println(" GenerateToken err:", err)
		c.JSON(500, gin.H{
			"status_code": -1,
		})
		return
	}
	// fmt.Println("token is")
	// fmt.Println(token)
	c.JSON(200, gin.H{
		"token":       token, //因为前端那里需要接收一个data*****************
		"user_id":     param.ID,
		"status_code": 0,
		"status_msg":  "sucess",
	})

}
