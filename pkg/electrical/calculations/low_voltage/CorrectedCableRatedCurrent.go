package low_voltage

import (
	"fmt"
	"math"
)

func CorrectedCableRatedCurrent(
	cableRatedCurrent float64,
	parallelCables int,
) (float64, error) {

	if math.IsNaN(cableRatedCurrent) || math.IsInf(cableRatedCurrent, 0) {
		return 0, fmt.Errorf("Corrente admissível do cabo inválida: valor não numérico ou infinito")
	}
	if cableRatedCurrent < 0 {
		return 0, fmt.Errorf(
			"Corrente admissível do cabo inválida: valor negativo (%.2f)",
			cableRatedCurrent,
		)
	}

	if parallelCables <= 0 {
		return 0, fmt.Errorf(
			"Quantidade de cabos por fase inválida: deve ser maior que zero",
		)
	}

	result := cableRatedCurrent * float64(parallelCables)

	return result, nil
}
