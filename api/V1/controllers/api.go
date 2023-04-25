package controllers

import "github.com/work/SRM/service"

type Api struct {
	AdminService      *service.AdminService
	ProgrammerService *service.ProgrammerService
	RegisterService   *service.RegisterService
}

func NewApi(admin *service.AdminService, programmer *service.ProgrammerService, register *service.RegisterService) *Api {
	return &Api{
		AdminService:      admin,
		ProgrammerService: programmer,
		RegisterService:   register,
	}
}
