package q1

import (
	"fmt"
)

func verificarElegebilidadeDesconto(valordaCompraAtual float64, historicodeCompras []float64) (float64, error) {
	if valordaCompraAtual <= 0 {
		return 0, fmt.Errorf("Valor da compra é invalido")
	}
	valorTotalHistorico := 0.0
	for _, valorDeCadaCompra := range historicodeCompras {
		valorTotalHistorico += valorDeCadaCompra
	}
	desconto := 0.0
	media := 0.0
	if valorTotalHistorico > 0 {
		media = valorTotalHistorico / float64(len(historicodeCompras))
	}

	if len(historicodeCompras) == 0 {
		desconto = valordaCompraAtual * 0.1
	} else if media > 1000 {
		desconto = valordaCompraAtual * 0.2
	} else if valorTotalHistorico > 1000 {
		desconto = valordaCompraAtual * 0.10
	} else if valorTotalHistorico <= 1000 && valorTotalHistorico > 500 {
		desconto = valordaCompraAtual * 0.05
	} else if valorTotalHistorico <= 500 {
		desconto = valordaCompraAtual * 0.02
	}

	return desconto, nil
}
