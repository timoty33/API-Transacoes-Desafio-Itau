package transacao

import "time"

type TransacaoDTORequest struct {
	Valor    float64   `json:"valor" binding:"required,gte=0"` // todo valor precisa ser maior que ou igual 0
	DataHora time.Time `json:"id" binding:"required"`
}

type TransacaoDTOResponse struct {
	Id        int       `json:"id"`
	Valor     float64   `json:"valor"`
	DataHora  time.Time `json:"data_hora"`
	Criado_em time.Time `json:"criado_em"`
}
