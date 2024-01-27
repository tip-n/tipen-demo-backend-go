package handler

import (
	"errors"
	"tipen-demo/pkg"
	"tipen-demo/repository"
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

func (h *Handler) LoginSeller(p LoginSellerParams) (jwt string, err error) {
	user, err := h.Repo.GetSellerByEmail(p.Email)
	if err != nil {
		return
	}

	isCorrect := pkg.CompareHash(user.Password, p.Password)
	if !isCorrect {
		err = errors.New("email atau password tidak sesuai")
		return
	}

	jwt, err = pkg.GenerateJWT(user.ID)
	if err != nil {
		return
	}

	err = h.Repo.InsertSellerLoginCount(int(user.ID))
	return
}

type SellerProfile struct {
	Storename string `json:"storename"`
	Email     string `json:"email"`
}

func (h *Handler) GetSellerProfile(ID int) (resp SellerProfile, err error) {
	seller, err := h.Repo.GetSellerProfile(ID)
	if err != nil {
		return
	}

	resp = SellerProfile{
		Storename: seller.Storename,
		Email:     seller.Email,
	}

	return
}
