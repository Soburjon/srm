package repo

import "github.com/work/SRM/api/V1/models"

type Register interface {
	Login(req models.RegisterRequest) (models.RegisterResponse, error)
}
