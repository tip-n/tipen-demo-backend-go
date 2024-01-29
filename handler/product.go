package handler

import (
	"tipen-demo/repository"

	"gorm.io/gorm"
)

type CreateProductParams struct {
	SellerID int
	Name     string
	Price    int
	Image    string
}

func (h *Handler) CreateProduct(p CreateProductParams) (ID int, err error) {
	ID, err = h.Repo.CreateProduct(repository.Products{
		SellerID: p.SellerID,
		Name:     p.Name,
		Price:    p.Price,
		Image:    p.Image,
	})
	return
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (h *Handler) GetProductByID(id int) (resp Product, err error) {
	product, err := h.Repo.GetProductByID(id)
	if err != nil {
		return
	}

	resp = Product{
		Name:  product.Name,
		Price: product.Price,
		Image: product.Image,
	}
	return
}

type ProductsParams struct {
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	SellerID int
}

type Products struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (h *Handler) GetProducts(p ProductsParams) (resp []Products, err error) {
	products, err := h.Repo.GetProductsBySellerID(repository.GetProductsBySellerIDParams{
		Limit:    p.Limit,
		Page:     p.Page,
		SellerID: p.SellerID,
	})
	for i := range products {
		resp = append(resp, Products{
			Name:  products[i].Name,
			Price: products[i].Price,
			Image: products[i].Image,
		})
	}
	return
}

type UpdateProductParams struct {
	ID    int
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (h *Handler) UpdateProduct(p UpdateProductParams) (err error) {
	product := repository.Products{
		Model: &gorm.Model{
			ID: uint(p.ID),
		},
		Name:  p.Name,
		Price: p.Price,
		Image: p.Image,
	}
	err = h.Repo.UpdateProduct(product)
	return
}
