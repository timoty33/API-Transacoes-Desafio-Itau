package utils

import (
	"time"
)

func validarTempo(old time.Time) bool {
	// A transação deve ter uma data no passado ou no presente.
	return !old.After(time.Now())
}

func validarValor(valor float64) bool {
	return valor >= 0
}

func ValidarTransacao(dataHora time.Time, valor float64) bool {
	if validarTempo(dataHora) && validarValor(valor) {
		return true
	}
	return false
}
