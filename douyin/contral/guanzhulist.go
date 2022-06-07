package contral

import (
	"douyin/global"
	"douyin/model"
	"douyin/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GuanzhuList(c *gin.Context) {
	userid := c.Query("user_id")
	theid, err := strconv.Atoi(userid)
	log.Println("test user_id:", theid)
	if err != nil {
		log.Println("theid strconv err:", err)
	}
	modelg := model.Guanzhu{Userid: int64(theid)}
	list, err := service.GetuserGuanzhuList(modelg, global.DbEngine)
	if err != nil {
		log.Println("guanzhu list get err:", err)
	}
	log.Println("guanzhu list :", list)
	c.JSON(200, gin.H{
		"status_code": 0,
		"user_list":   list,
	})
}

func FensiList(c *gin.Context) {
	userid := c.Query("user_id")
	theid, err := strconv.Atoi(userid)
	if err != nil {
		log.Println("theid strconv err:", err)
	}
	modelg := model.Guanzhu{Touserid: int64(theid)}
	list, err := service.GetuserFensiList(modelg, global.DbEngine)
	if err != nil {
		log.Println("guanzhu list get err:", err)
	}
	log.Println("fensi list :", list)
	c.JSON(200, gin.H{
		"status_code": 0,
		"user_list":   list,
	})
}
