package data

import "time"

type Transacao struct {
	Valor    float64   `json:"valor" binding:"required"`
	DataHora time.Time `json:"dataHora" binding:"required"`
}

var Transacoes []Transacao
