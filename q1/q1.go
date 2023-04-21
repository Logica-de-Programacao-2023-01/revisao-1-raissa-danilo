package q1

import (
	"errors"
)

func calcular(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase <= 0 {
		return 0, errors.New("valor da compra inválido")
	}

	var totalPurchases float64
	for _, purchase := range purchaseHistory {
		totalPurchases += purchase
	}

	var discount float64
	if totalPurchases > 1000 {
		discount = 0.1
	} else if totalPurchases > 500 {
		discount = 0.05
	} else {
		discount = 0.02
	}

	if len(purchaseHistory) == 0 {
		discount = 0.1
	}

	var averagePurchase float64
	if len(purchaseHistory) > 0 {
		averagePurchase = totalPurchases / float64(len(purchaseHistory))
	}

	if averagePurchase > 1000 {
		discount = 0.2
	}

	return currentPurchase * discount, nil
}
