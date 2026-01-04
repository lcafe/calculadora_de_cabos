package pkg

import (
	"fmt"
	"math"
)

func BreakerConventionalTripCurrent(
	breakerRatedCurrent float64,
) (float64, error) {

	if math.IsNaN(breakerRatedCurrent) || math.IsInf(breakerRatedCurrent, 0) {
		return 0, fmt.Errorf("Corrente nominal do disjuntor inválida: valor não numérico ou infinito")
	}
	if breakerRatedCurrent < 0 {
		return 0, fmt.Errorf(
			"Corrente nominal do disjuntor inválida: valor negativo (%.2f)",
			breakerRatedCurrent,
		)
	}

	result := 1.30 * breakerRatedCurrent

	return result, nil
}
