package models

type House struct {
	ID      int    `json: "id"`
	Title   string `json:"title`
	Content string `json:"content"`
	Address string `json:"address"`
}
