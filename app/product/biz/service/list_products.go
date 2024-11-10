package service

import (
	"Gomall/app/product/biz/dal/mysql"
	"Gomall/app/product/biz/model"
	product "Gomall/rpc_gen/kitex_gen/product"
	"context"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductReq) (resp *product.ListProductResp, err error) {
	// Finish your business logic.
	CategoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := CategoryQuery.GetProductsByCategoryName(req.CategoryName)
	resp = &product.ListProductResp{}
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
			})
		}
	}
	return resp, nil
}
