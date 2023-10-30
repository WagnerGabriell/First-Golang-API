package entities

import "github.com/google/uuid"

type Produto struct {
	ID         string  `json:"id"`
	Nome       string  `json:"nome"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

func NewProduto() *Produto {
	prod := Produto{
		ID: uuid.New().String(),
	}
	return &prod
}
