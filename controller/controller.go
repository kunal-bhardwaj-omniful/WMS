package controller

import (
	"wms/service"
)

type Controller struct {
	service *service.Service
}

// NewUserController is the constructor for UserController, it takes a repository as dependency.
func NewUserController(s *service.Service) *Controller {
	return &Controller{
		service: s,
	}
}
