package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Sellers struct {
	*gorm.Model
	Storename string
	Password  string
	Email     string
}

type SellerLogins struct {
	*gorm.Model
	SellerId string
}

func (r *Repository) RegisterSeller(p Sellers) (ID int, err error) {
	if isExist := r.CheckSellerExistByEmail(p.Email); isExist {
		err = errors.New("seller already registered")
		return
	}

	err = r.Db.Create(&p).Error
	if err != nil {
		return
	}
	ID = int(p.ID)
	return
}

func (r *Repository) InsertSellerLoginCount(ID int) (err error) {
	err = r.Db.Model(
		&SellerLogins{},
	).Create(map[string]interface{}{"seller_id": ID}).Error
	return
}

func (r *Repository) CheckSellerExistByEmail(email string) (isExist bool) {
	ID := 0
	isExist = false
	r.Db.Model(&Sellers{}).
		Select("id").
		Where(&Sellers{Email: email}).
		First(&ID)
	if ID != 0 {
		isExist = true
	}
	return
}

func (r *Repository) GetSellerByEmail(email string) (
	resp Sellers,
	err error,
) {
	err = r.Db.Model(&Sellers{}).
		Where(&Sellers{Email: email}).
		First(&resp).Error
	return
}

func (r *Repository) GetSellerByID(ID int) (
	resp Sellers,
	err error,
) {
	err = r.Db.Model(&Sellers{}).
		First(&resp, ID).Error
	return
}

func (r *Repository) UpdateSeller(p Sellers) (err error) {
	seller, err := r.GetSellerByID(int(p.ID))
	if err != nil {
		return
	}

	if p.Storename != "" {
		seller.Storename = p.Storename
	}

	err = r.Db.Save(&seller).Error
	return
}
