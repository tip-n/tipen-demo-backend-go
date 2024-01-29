package repository

import (
	"errors"

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
	if isExist := r.CheckUserExistByEmail(p.Email); isExist {
		err = errors.New("user already registered")
		return
	}

	err = r.Db.Create(&p).Error
	if err != nil {
		return
	}
	ID = int(p.ID)
	return
}

func (r *Repository) InsertUserLoginCount(ID int) (err error) {
	err = r.Db.Model(
		&UserLogins{},
	).Create(map[string]interface{}{"user_id": ID}).Error
	return
}

func (r *Repository) CheckUserExistByEmail(email string) (isExist bool) {
	ID := 0
	isExist = true
	err := r.Db.Model(&Users{}).
		Select("id").
		Where(&Users{Email: email}).
		First(&ID).Error
	if ID == 0 || err == gorm.ErrRecordNotFound {
		isExist = false
		return
	}
	return
}

func (r *Repository) GetUserByEmail(email string) (
	resp Users,
	err error,
) {
	err = r.Db.Model(&Users{}).
		Where(&Users{Email: email}).
		First(&resp).Error
	return
}

func (r *Repository) GetUserByID(ID int) (
	resp Users,
	err error,
) {
	err = r.Db.Model(&Users{}).
		First(&resp, ID).Error
	return
}

func (r *Repository) UpdateUser(p Users) (err error) {
	user, err := r.GetUserByID(int(p.ID))
	if err != nil {
		return
	}

	if p.Firstname != "" {
		user.Firstname = p.Firstname
	}

	err = r.Db.Save(&user).Error
	return
}
