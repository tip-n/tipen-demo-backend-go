package repository

import (
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
	err = r.Db.Create(&p).Error
	if err != nil {
		return
	}
	ID = int(p.ID)
	return
}

type GetSellerPasswordByEmailResponse struct {
	ID       int64
	Password string
}

func (r *Repository) GetSellerByEmail(email string) (
	resp GetSellerPasswordByEmailResponse,
	err error,
) {
	err = r.Db.Model(&Sellers{}).
		Select("id", "password").
		Where(&Sellers{Email: email}).
		First(&resp).Error
	return
}

func (r *Repository) InsertSellerLoginCount(ID int) (err error) {
	err = r.Db.Model(
		&SellerLogins{},
	).Create(map[string]interface{}{"seller_id": ID}).Error
	return
}

type GetSellerProfileResponse struct {
	Storename string
	Email     string
}

func (r *Repository) GetSellerProfile(ID int) (
	resp GetSellerProfileResponse,
	err error,
) {
	err = r.Db.Model(&Sellers{}).
		Select("storename", "email").
		First(&resp, ID).Error
	return
}
