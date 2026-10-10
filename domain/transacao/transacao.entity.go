package transacao

import "time"

type transacao struct {
	Id        int
	Valor     float64
	DataHora  time.Time
	Criado_em time.Time 
}
