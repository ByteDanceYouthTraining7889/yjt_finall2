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

func Guanzhu(c *gin.Context) {
	param := model.Guanzhu{}
	token := c.Query("token")
	claims, _ := pkg.ParseToken(token)
	username := claims.Id
	userpass := claims.Secret
	fmt.Println("test username:", username)

	modelu := model.User{Username: username, Password: userpass}
	user, _ := service.GetUser(modelu, global.DbEngine)

	valid, errs := pkg.BindAndValid(c, &param)
	fmt.Println("the param:", param)
	if !valid {
		log.Println("CreatUserInDenglu bindandvalid err:", errs)
		// c.JSON(500, gin.H{
		// 	"status_code": -1,
		// })
		// return
	}
	param.Userid = user.ID
	fmt.Println("test use_id:", param.Userid)
	if param.Actiontype == 1 {
		err := service.Createguanzhu(param, global.DbEngine)
		if err != nil {
			log.Println("create guanzhu bindandvalid err:", errs)
			// c.JSON(500, gin.H{
			// 	"status_code": -1,
			// })
			// return
		}
		modelu := model.User{ID: param.Touserid}
		user, err := service.GetUserbyId(modelu, global.DbEngine)
		if err != nil {
			log.Println("getguanzhu userby id  err:", err)
			// c.JSON(500, gin.H{
			// 	"status_code": -1,
			// })
			// return
		}

		user.FollowCount += 1
		user.Is_follow = true
		err = service.UpdateUserguanzhu(user, global.DbEngine)
		if err != nil {
			log.Println("updatecomment    err:", err)
			// c.JSON(500, gin.H{
			// 	"status_code": -1,
			// })
			// return
		}
		c.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "sucess",
		})

	}
	if param.Actiontype == 2 {
		err := service.DeleteGuanzhu(param, global.DbEngine)
		modelu := model.User{ID: param.Touserid}
		user, err := service.GetUserbyId(modelu, global.DbEngine)
		if err != nil {
			log.Println("getguanzhu userby id  err:", err)
			// c.JSON(500, gin.H{
			// 	"status_code": -1,
			// })
			// return
		}

		user.FollowCount += -1
		user.Is_follow = false
		err = service.UpdateUserguanzhu(user, global.DbEngine)
		if err != nil {
			log.Println("updatecomment    err:", err)
			// c.JSON(500, gin.H{
			// 	"status_code": -1,
			// })
			// return
		}
		if err != nil {
			log.Println("create guanzhu bindandvalid err:", errs)
			// c.JSON(500, gin.H{
			// 	"status_code": -1,
			// })
			// return
		}
		c.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "sucess",
		})

	}
}
