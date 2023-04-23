package q2

import (
	"errors"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	caracteresEspeciais := []string{".", ",", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	for _, caracteres := range caracteresEspeciais {
		text = strings.ReplaceAll(text, caracteres, "")
	}

	palavras := strings.Fields(text)
	if len(palavras) == 0 {
		return 0, errors.New("texto está vazio")
	}

	var numeroTotalDeLetras = 0

	for _, palavra := range palavras {
		numeroTotalDeLetras += len(palavra)
	}

	return float64(numeroTotalDeLetras) / float64(len(palavras)), nil
}
