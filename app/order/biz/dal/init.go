package dal

import (
	"Gomall/app/order/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
