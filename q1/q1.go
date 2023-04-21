package q1

import (
	"errors"
)

func verificarElegebilidadeDesconto(valordaCompraAtual float64, historicodeCompras []float64) (float64, error) {
	if valordaCompraAtual <= 0 {
		return 0, errors.New("Valor da compra é invalido")
	}
	valorTotalHistorico := 0.0
	for _, valorDeCadaCompra := range historicodeCompras {
		valorTotalHistorico += valorDeCadaCompra
	}
	media := 0.0
	if valorTotalHistorico > 0 {
		media = valorTotalHistorico / float64(len(historicodeCompras))
	}

	var discount float64
	if valorTotalHistorico > 1000 {
		discount = 0.1
	} else if valorTotalHistorico > 500 {
		discount = 0.05
	} else {
		discount = 0.02
	}

	if len(historicodeCompras) == 0 {
		discount = 0.1
	}
	
	if media > 1000 {
		discount = 0.2
	}
	
	return discount * valordaCompraAtual, nil
}
