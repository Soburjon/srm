package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/work/SRM/api/V1/models"
	"github.com/work/SRM/storage"
)

type RegisterService struct {
	storage storage.IStorage
}

func NewRegisterService(db *sqlx.DB) *RegisterService {
	return &RegisterService{
		storage: storage.New(db),
	}
}

func (r *RegisterService) Login(req models.RegisterRequest) (models.RegisterResponse, error) {
	res, err := r.storage.Register().Login(req)
	if err != nil {
		return models.RegisterResponse{}, err
	}
	return res, nil
}
