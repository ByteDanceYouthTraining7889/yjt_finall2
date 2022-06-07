package main

import (
	"douyin/contral"
	"douyin/global"
	"douyin/model"
	"douyin/setting"
)

// docker run -it --network host --rm mysql mysql -h127.0.0.1 -P13306 --default-character-set=utf8mb4 -uroot -p
func init() {
	dbs := setting.SetDB()
	global.DbEngine, _ = model.NewDBEngine(dbs)
	global.JWTSettings = setting.SetJWT()
	global.FileSave = setting.SetUploadFile()
}
func main() {
	r := contral.NewRouter()
	r.Run()
}
