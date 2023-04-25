package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/work/SRM/api/V1/models"
	"github.com/work/SRM/pkg/sqls"
	"github.com/work/SRM/storage/repo"
)

type registerRepo struct {
	db *sqlx.DB
}

func NewRegisterRepo(db *sqlx.DB) repo.Register {
	return &registerRepo{
		db: db,
	}
}

func (r *registerRepo) Login(req models.RegisterRequest) (res models.RegisterResponse, err error) {
	if err = r.db.QueryRow(sqls.Login, req.PhoneNumber, req.Password).Scan(&res.UserID, &res.Role); err != nil {
		return res, err
	}
	return res, nil
}
