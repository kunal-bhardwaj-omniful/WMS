package service

import (
	"context"
	"wms/domain"
	"wms/repo"
)

type Service interface {
	FetchHubs(ctx context.Context) []domain.Hub
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

func (s *service) FetchHubs(ctx context.Context) []domain.Hub {
	return s.repo.GetAllHubs(ctx)
}
