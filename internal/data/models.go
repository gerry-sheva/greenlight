package data

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	errRecordNotFound = errors.New("Record not found")
)

type Models struct {
	Movies MovieModel
}

func NewModels(db *pgxpool.Pool) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
