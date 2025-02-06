package service

import (
	"wms/repo"
)

type Service interface {
}

type service struct {
	repo repo.Repository
}

// NewService is the constructor function to create a new instance of ConcreteService.
func NewService(r repo.Repository) Service {
	return &service{
		repo: r,
	}
}
