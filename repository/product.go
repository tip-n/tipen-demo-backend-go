package repository

import (
	"gorm.io/gorm"
)

type Products struct {
	*gorm.Model
	SellerID int
	Name     string
	Price    int
	Image    string
}

func (r *Repository) CreateProduct(p Products) (ID int, err error) {
	err = r.Db.Create(&p).Error
	if err != nil {
		return
	}
	ID = int(p.ID)
	return
}

func (r *Repository) GetProductByID(ID int) (
	resp Products,
	err error,
) {
	err = r.Db.Model(&Products{}).
		First(&resp, ID).Error
	return
}

type GetProductsBySellerIDParams struct {
	Limit    int
	Page     int
	SellerID int
}

func (r *Repository) GetProductsBySellerID(p GetProductsBySellerIDParams) (
	resp []Products,
	err error,
) {
	err = r.Db.Find(&resp).
		Where(&Products{SellerID: p.SellerID}).
		Limit(p.Limit).Offset(p.Page - 1).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
		resp = []Products{}
	}
	return
}

func (r *Repository) CheckProductExistByName(name string) (isExist bool) {
	ID := 0
	isExist = false
	r.Db.Model(&Products{}).
		Select("id").
		Where(&Products{Name: name}).
		First(&ID)
	if ID != 0 {
		isExist = true
	}
	return
}

func (r *Repository) UpdateProduct(p Products) (err error) {
	product, err := r.GetProductByID(int(p.ID))
	if err != nil {
		return
	}

	if p.Name != "" {
		product.Name = p.Name
	}
	if p.Image != "" {
		// usually you will need to delete the file too
		product.Image = p.Image
	}
	if p.Price > 0 {
		product.Price = p.Price
	}

	err = r.Db.Save(&product).Error
	return
}
