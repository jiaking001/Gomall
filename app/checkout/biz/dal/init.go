package dal

import (
	"Gomall/app/checkout/biz/dal/mysql"
	"Gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
