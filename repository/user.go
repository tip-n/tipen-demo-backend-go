package repository

import (
	"gorm.io/gorm"
)

type Users struct {
	*gorm.Model
	Firstname string
	Lastname  string
	Password  string
	Email     string
}

type UserLogins struct {
	*gorm.Model
	UserId string
}

func (r *Repository) RegisterUser(p Users) (ID int, err error) {
	err = r.Db.Create(&p).Error
	if err != nil {
		return
	}
	ID = int(p.ID)
	return
}

type GetUserPasswordByEmailResponse struct {
	ID       int64
	Password string
}

func (r *Repository) GetUserByEmail(email string) (
	resp GetUserPasswordByEmailResponse,
	err error,
) {
	err = r.Db.Model(&Users{}).
		Select("id", "password").
		Where(&Users{Email: email}).
		First(&resp).Error
	return
}

func (r *Repository) InsertUserLoginCount(ID int) (err error) {
	err = r.Db.Model(
		&UserLogins{},
	).Create(map[string]interface{}{"user_id": ID}).Error
	return
}

type GetUserProfileResponse struct {
	Firstname string
	Lastname  string
	Email     string
}

func (r *Repository) GetUserProfile(ID int) (
	resp GetUserProfileResponse,
	err error,
) {
	err = r.Db.Model(&Users{}).
		Select("firstname", "lastname", "email").
		First(&resp, ID).Error
	return
}
