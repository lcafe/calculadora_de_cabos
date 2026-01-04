package low_voltage

import (
	"fmt"
	"math"
)

func CableThermalLimitCurrent(
	correctedCableCurrent float64,
) (float64, error) {

	if math.IsNaN(correctedCableCurrent) || math.IsInf(correctedCableCurrent, 0) {
		return 0, fmt.Errorf("Corrente admissível corrigida inválida: valor não numérico ou infinito")
	}
	if correctedCableCurrent < 0 {
		return 0, fmt.Errorf(
			"Corrente admissível corrigida inválida: valor negativo (%.2f)",
			correctedCableCurrent,
		)
	}

	result := 1.45 * correctedCableCurrent

	return result, nil
}
