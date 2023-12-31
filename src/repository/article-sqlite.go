package repository

import (
	model "datareporter/model"
)

type ArticleSqlite struct{}

func (repository *ArticleSqlite) GetDbEngine() string {
	return "sqlite"
}

func (repository *ArticleSqlite) Get() (*[]model.Article, error) {
	db, err := Con.OpenCon(repository.GetDbEngine())
	if err != nil {
		return nil, err
	}

	defer db.Close()
	if db.Ping(); err != nil {
		return nil, err
	}

	query := `SELECT article_id AS id,
                     title,
                     body,
                     author_code,
                     is_active,
                     created_on,
                     updated_on
                FROM articles
               WHERE is_active = true
            ORDER BY title`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var articles []model.Article
	for rows.Next() {
		var model model.Article
		if err = rows.Scan(
			&model.Id,
			&model.Title,
			&model.Body,
			&model.AuthorCode,
			&model.Active,
			&model.Created,
			&model.Updated,
		); err != nil {
			return nil, err
		}

		articles = append(articles, model)
	}

	return &articles, nil
}
