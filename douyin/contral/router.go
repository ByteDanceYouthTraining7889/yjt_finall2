package contral

import (
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./public") //将url上对应的static 改为public
	r.POST("/douyin/user/register/", CreatUserInDenglu)
	r.GET("/douyin/feed", GetVideoFeed)
	r.POST("/douyin/user/login/", UserInlogin)
	douyin := r.Group("/douyin")
	douyin.Use(middleware.JWT())
	douyin.GET("user/", GetUserMsg)
	douyin.POST("publish/action/", FabuVideo)
	douyin.GET("publish/list/", PublishList)
	// extra apis - I
	douyin.POST("favorite/action/", UpdateZan)
	douyin.GET("favorite/list/", DianzanList)
	douyin.POST("comment/action/", CreateComment)
	douyin.GET("comment/list/", CommentList)
	//-II
	douyin.POST("relation/action/", Guanzhu)
	douyin.GET("relation/follow/list/", GuanzhuList)
	douyin.GET("relation/follower/list/", FensiList)
	return r
}
