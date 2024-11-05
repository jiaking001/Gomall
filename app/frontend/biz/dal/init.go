package dal

import (
	"Gomall/app/frontend/biz/dal/mysql"
	"Gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
