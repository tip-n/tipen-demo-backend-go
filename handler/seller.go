package handler

import (
	"errors"
	"time"
	"tipen-demo/pkg"
	"tipen-demo/repository"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type RegisterSellerParams struct {
	Storename string
	Password  string
	Email     string
}

func (h *Handler) RegisterSeller(p RegisterSellerParams) (ID int, err error) {
	hashedPassword, err := pkg.HashPassword(p.Password)
	if err != nil {
		return
	}

	ID, err = h.Repo.RegisterSeller(repository.Sellers{
		Storename: p.Storename,
		Password:  hashedPassword,
		Email:     p.Email,
	})
	return
}

type LoginSellerParams struct {
	Email    string
	Password string
}

func (h *Handler) LoginSeller(p LoginSellerParams) (token string, err error) {
	seller, err := h.Repo.GetSellerByEmail(p.Email)
	if err != nil {
		return
	}

	isCorrect := pkg.CompareHash(seller.Password, p.Password)
	if !isCorrect {
		err = errors.New("email atau password tidak sesuai")
		return
	}

	t := time.Now()
	claims := jwt.MapClaims{
		"seller_id": seller.ID,
		"iat":       t.Unix(),
		"exp":       t.Add(time.Hour * 24).Unix(),
	}
	token, err = pkg.GenerateJWT(claims)
	if err != nil {
		return
	}

	err = h.Repo.InsertSellerLoginCount(int(seller.ID))
	return
}

type SellerProfile struct {
	Storename string `json:"storename"`
	Email     string `json:"email"`
}

func (h *Handler) GetSellerProfile(ID int) (resp SellerProfile, err error) {
	seller, err := h.Repo.GetSellerByID(ID)
	if err != nil {
		return
	}

	resp = SellerProfile{
		Storename: seller.Storename,
		Email:     seller.Email,
	}

	return
}

type UpdateSellerParams struct {
	ID        int
	Storename string `json:"storename"`
}

func (h *Handler) UpdateSeller(p UpdateSellerParams) (err error) {
	seller := repository.Sellers{
		Model: &gorm.Model{
			ID: uint(p.ID),
		},
		Storename: p.Storename,
	}
	err = h.Repo.UpdateSeller(seller)
	return
}
