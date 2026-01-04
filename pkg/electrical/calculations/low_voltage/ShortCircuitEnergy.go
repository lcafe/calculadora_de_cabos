package low_voltage

import (
	"fmt"
	"math"
)

func ShortCircuitEnergy(
	current float64,
	time float64,
) (float64, error) {

	if math.IsNaN(current) || math.IsInf(current, 0) {
		return 0, fmt.Errorf("Corrente inválida: valor não numérico ou infinito")
	}
	if current < 0 {
		return 0, fmt.Errorf("Corrente inválida: deve ser maior ou igual a zero (%.2f)", current)
	}

	if math.IsNaN(time) || math.IsInf(time, 0) {
		return 0, fmt.Errorf("Tempo inválido: valor não numérico ou infinito")
	}
	if time < 0 {
		return 0, fmt.Errorf("Tempo inválido: deve ser maior ou igual a zero (%.6f)", time)
	}

	result := math.Pow(current, 2) * time
	return result, nil
}
