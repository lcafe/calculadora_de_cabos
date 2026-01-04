package pkg

import (
	"fmt"
	"math"
)

func CheckVoltageDrop(
	deltaPercent float64,
	limitPercent float64,
) (bool, error) {

	if math.IsNaN(deltaPercent) || math.IsInf(deltaPercent, 0) {
		return false, fmt.Errorf("Queda de tensão inválida: valor não numérico ou infinito")
	}
	if deltaPercent < 0 {
		return false, fmt.Errorf(
			"Queda de tensão inválida: valor negativo (%.2f)",
			deltaPercent,
		)
	}

	if math.IsNaN(limitPercent) || math.IsInf(limitPercent, 0) {
		return false, fmt.Errorf("Limite de queda de tensão inválido: valor não numérico ou infinito")
	}
	if limitPercent < 0 {
		return false, fmt.Errorf(
			"Limite de queda de tensão inválido: valor negativo (%.2f)",
			limitPercent,
		)
	}

	var result bool

	if deltaPercent <= limitPercent {
		result = true
	} else {
		result = false
	}

	return result, nil
}
