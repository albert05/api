package services

import (
	"github.com/astaxie/beego"
	"api/util/mysql"
	_ "github.com/astaxie/beego/session/redis"
	"api/util/redis"
	"api/common"
)

func ConfigInit() {
	// init mysql
	dbDsn := beego.AppConfig.String("dbconfig::dsn")
	mysql.Init(dbDsn)

	host := beego.AppConfig.String("redisconfig::host")
	password := beego.AppConfig.String("redisconfig::password")
	db := beego.AppConfig.String("redisconfig::db")

	// init redis
	redis.Init(host, password, common.Str2Int(db))
}
