package controller

import (
	"wms/service"
)

type Controller struct {
	service *service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{
		service: s,
	}
}
