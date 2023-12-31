package repository

import (
	"datareporter/model"
	"datareporter/repository/connection"
)

var Con = connection.Generic{}

type IRepository interface {
	Get() (*[]model.Article, error)
	GetDbEngine() (dbEngine string)
}
