// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"Gomall/app/frontend/biz/router/auth"
	category "Gomall/app/frontend/biz/router/category"
	home "Gomall/app/frontend/biz/router/home"
	product "Gomall/app/frontend/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	category.Register(r)

	product.Register(r)

	auth.Register(r)

	home.Register(r)
}
