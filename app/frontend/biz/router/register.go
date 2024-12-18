// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"Gomall/app/frontend/biz/router/auth"
	cart "Gomall/app/frontend/biz/router/cart"
	category "Gomall/app/frontend/biz/router/category"
	frontend_checkout "Gomall/app/frontend/biz/router/checkout"
	home "Gomall/app/frontend/biz/router/home"
	order "Gomall/app/frontend/biz/router/order"
	product "Gomall/app/frontend/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	order.Register(r)

	frontend_checkout.Register(r)

	cart.Register(r)

	category.Register(r)

	product.Register(r)

	auth.Register(r)

	home.Register(r)
}
