package repository

import "gorm.io/gorm"

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

type GetPasswordByEmailResponse struct {
	ID       int64
	Password string
}

func (r *Repository) GetUserByEmail(email string) (
	resp GetPasswordByEmailResponse,
	err error,
) {
	err = r.Db.Model(&Users{}).Select("id", "password").First(&resp).Error
	return
}

func (r *Repository) InsertUserLoginCount(ID int) (err error) {
	err = r.Db.Model(
		&UserLogins{},
	).Create(map[string]interface{}{"user_id": ID}).Error
	return
}
