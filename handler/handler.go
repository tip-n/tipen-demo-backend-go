package handler

import (
	"tipen-demo/repository"
)

type Handler struct {
	Repo *repository.Repository
	// add later : redis to cache
}

type HandlerParams struct {
	Repository *repository.Repository
}

func NewHandler(p HandlerParams) *Handler {
	return &Handler{
		Repo: p.Repository,
	}
}
