package model

type QualquerCoisa struct {
	Id    uint64               `json:"id"`
	Code  string               `json:"code"`
	I18ns *[]QualquerCoisaI18n `json:"i18n"`
}

type QualquerCoisaI18n struct {
	Name string `json:"name"`
	Lang string `json:"lang"`
}
