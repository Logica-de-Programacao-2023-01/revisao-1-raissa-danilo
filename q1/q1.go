package q1

import (
	"fmt"
)

func CalculateDiscount(valorCompraAtual float64, historicoCompras []float64) (float64, error) {
	if valorCompraAtual <= 0 {
		return 0, fmt.Errorf("Valor da compra invalido")
	}

	valorTotalHistorico := 0.0
	for _, valorDeCadaCompra := range historicoCompras {
		valorTotalHistorico += valorDeCadaCompra
	}

	desconto := 0.0
	media := 0.0

	if valorTotalHistorico > 0 {
		media = valorTotalHistorico / float64(len(historicoCompras))
	}

	if len(historicoCompras) == 0 {
		desconto = valorCompraAtual * 0.1
	} else if media > 1000 {
		desconto = valorCompraAtual * 0.2
	} else if valorTotalHistorico > 1000 {
		desconto = valorCompraAtual * 0.1
	} else if valorTotalHistorico <= 1000 && valorTotalHistorico > 500 {
		desconto = valorCompraAtual * 0.05
	} else if valorTotalHistorico <= 500 {
		desconto = valorCompraAtual * 0.02
	}

	return desconto, nil
}
