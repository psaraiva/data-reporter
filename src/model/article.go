package model

type Article struct {
	Id         uint64 `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	AuthorCode string `json:"author-code"`
	Version    uint8  `json:"version"`
	Active     bool   `json:"active"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}
