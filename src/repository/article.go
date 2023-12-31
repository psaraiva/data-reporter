package repository

import (
	"datareporter/config"
	"datareporter/model"
	"errors"
)

type Article struct{}

var env = config.AppConfig{}

func (a *Article) loadRepository() (IRepository, error) {
	dbEngine, err := env.LoadDbEngine()
	if err != nil {
		return nil, err
	}

	return a.getRepository(dbEngine)
}

func (a *Article) getRepository(dbEngine string) (IRepository, error) {
	switch dbEngine {
	case "postgres":
		return &ArticlePostgres{}, nil
	case "sqlite":
		return &ArticleSqlite{}, nil
	default:
		return nil, errors.New("invalid database engine")
	}
}

func (a *Article) Get() (*[]model.Article, error) {
	reporitory, err := a.loadRepository()
	if err != nil {
		return nil, err
	}

	return reporitory.Get()
}
