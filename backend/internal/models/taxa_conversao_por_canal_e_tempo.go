package models

import "time"

type TaxaConversaoPorCanalETempo struct {
	Data          time.Time `json:"data"`
	Canal         string    `json:"canal"`
	TaxaConversao float64   `json:"taxa_conversao"`
}
