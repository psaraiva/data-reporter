package controller

import (
	"log"
	"net/http"

	"datareporter/model"
	"datareporter/repository"

	"github.com/go-chi/render"
)

type Article struct{}

var articleRepository = repository.Article{}

var envelope = Envelope[model.Article]{
	IsErro:   false,
	ErroCode: "",
	Message:  "",
}

func (controller *Article) List(w http.ResponseWriter, r *http.Request) {
	articles, err := articleRepository.Get()
	if err != nil {
		envelope.IsErro = true
		envelope.ErroCode = "ERR:799"
		envelope.Message = "Error on list Article"

		log.Fatal(err)

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, envelope)
		return
	}

	envelope.Data = articles
	render.Status(r, http.StatusOK)
	render.JSON(w, r, envelope)
}
