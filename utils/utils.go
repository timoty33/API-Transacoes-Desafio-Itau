package utils

import (
	"fmt"
	"time"
	"transactionsApi/data"
)

// --- Funções de validação ---
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

// --- Funções de Estatísticas ---
func UltimasTransacoes(lista []data.Transacao, duracao time.Duration) []data.Transacao {
	var ultimos []data.Transacao
	agora := time.Now()
	limiteTempo := agora.Add(-duracao)

	for _, transacao := range lista {
		// Verifica se a transação está dentro do período dos últimos 60 segundos
		if transacao.DataHora.After(limiteTempo) {
			ultimos = append(ultimos, transacao)
		}
	}

	// Retorna o slice de todas as transações dos últimos 60 segundos
	return ultimos
}

func Soma(lista []data.Transacao) float64 { // Soma não gera erro, mas se futuramente modificar para ter, é aqui
	var total float64
	for _, l := range lista {
		total += l.Valor
	}
	return total
}

func Media(lista []data.Transacao) (float64, error) {
	if len(lista) == 0 {
		return 0, fmt.Errorf("não é possível calcular a média de uma lista vazia")
	}
	return Soma(lista) / float64(len(lista)), nil
}

func MinMax(lista []data.Transacao) (float64, float64, error) {
	if len(lista) == 0 {
		return 0, 0, fmt.Errorf("não é possível calcular min/max de uma lista vazia")
	}

	min := lista[0].Valor
	max := lista[0].Valor

	for _, transacao := range lista {
		if transacao.Valor < min {
			min = transacao.Valor
		}
		if transacao.Valor > max {
			max = transacao.Valor
		}
	}
	return min, max, nil
}
