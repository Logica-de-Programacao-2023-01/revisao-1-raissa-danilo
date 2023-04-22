package q1

import (
	"fmt"
)

func CalculateDiscount(valorCompraAtual float64, historicoCompras []float64) (float64, error) {
	if valorCompraAtual <= 0 {
		return 0, fmt.Errorf("Valor da compra invalido.")
	}

	media := 0.0
	desconto := 0.0

	valorHistoricoTotal := 0.0
	for _, valorCadaCompra := range historicoCompras {
		valorHistoricoTotal += valorCadaCompra
	}

	if valorHistoricoTotal > 0 {
		media = valorHistoricoTotal / float64(len(historicoCompras))
	}

	if len(historicoCompras) == 0 {
		desconto = valorCompraAtual * 0.1
	} else if media > 1000 {
		desconto = valorCompraAtual * 0.2
	} else if valorHistoricoTotal > 1000 {
		desconto = valorCompraAtual * 0.1
	} else if valorHistoricoTotal <= 1000 && valorHistoricoTotal > 500 {
		desconto = valorCompraAtual * 0.05
	} else if valorHistoricoTotal <= 500 {
		desconto = valorCompraAtual * 0.02
	}

	return desconto, nil
}
