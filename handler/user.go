package handler

import (
	"errors"
	"tipen-demo/pkg"
	"tipen-demo/repository"
)

type RegisterUserParams struct {
	Firstname string
	Lastname  string
	Password  string
	Email     string
}

func (h *Handler) RegisterUser(p RegisterUserParams) (ID int, err error) {
	hashedPassword, err := pkg.HashPassword(p.Password)
	if err != nil {
		return
	}

	ID, err = h.Repo.RegisterUser(repository.Users{
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Password:  hashedPassword,
		Email:     p.Email,
	})
	return
}

type LoginUserParams struct {
	Email    string
	Password string
}

func (h *Handler) LoginUser(p LoginUserParams) (jwt string, err error) {
	user, err := h.Repo.GetUserByEmail(p.Email)
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

	err = h.Repo.InsertUserLoginCount(int(user.ID))
	return
}
