package pkg

import (
	"fmt"
	"math"
)

func CheckShortCircuit(
	shortCircuitEnergy float64,
	cableThermalCapacity float64,
) (bool, error) {

	if math.IsNaN(shortCircuitEnergy) || math.IsInf(shortCircuitEnergy, 0) {
		return false, fmt.Errorf("Energia de curto-circuito inválida: valor não numérico ou infinito")
	}
	if shortCircuitEnergy < 0 {
		return false, fmt.Errorf(
			"Energia de curto-circuito inválida: valor negativo (%.2f)",
			shortCircuitEnergy,
		)
	}

	if math.IsNaN(cableThermalCapacity) || math.IsInf(cableThermalCapacity, 0) {
		return false, fmt.Errorf("Capacidade térmica do cabo inválida: valor não numérico ou infinito")
	}
	if cableThermalCapacity < 0 {
		return false, fmt.Errorf(
			"Capacidade térmica do cabo inválida: valor negativo (%.2f)",
			cableThermalCapacity,
		)
	}

	var result bool

	if shortCircuitEnergy <= cableThermalCapacity {
		result = true
	} else {
		result = false
	}

	return result, nil
}
