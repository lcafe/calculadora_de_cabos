package low_voltage

import (
	"fmt"
	"math"
)

func VoltageDrop(
	current float64,
	length float64,
	resistance float64,
	reactance float64,
	powerFactor float64,
) (float64, error) {

	if math.IsNaN(current) || math.IsInf(current, 0) {
		return 0, fmt.Errorf("Corrente inválida: valor não numérico ou infinito")
	}
	if current < 0 {
		return 0, fmt.Errorf("Corrente inválida: deve ser não negativa (%.2f)", current)
	}

	if math.IsNaN(length) || math.IsInf(length, 0) {
		return 0, fmt.Errorf("Comprimento inválido: valor não numérico ou infinito")
	}
	if length < 0 {
		return 0, fmt.Errorf("Comprimento inválido: deve ser não negativo (%.2f)", length)
	}

	if math.IsNaN(resistance) || math.IsInf(resistance, 0) {
		return 0, fmt.Errorf("Resistência inválida: valor não numérico ou infinito")
	}
	if resistance < 0 {
		return 0, fmt.Errorf("Resistência inválida: deve ser não negativa (%.6f)", resistance)
	}

	if math.IsNaN(reactance) || math.IsInf(reactance, 0) {
		return 0, fmt.Errorf("Reatância inválida: valor não numérico ou infinito")
	}
	if reactance < 0 {
		return 0, fmt.Errorf("Reatância inválida: deve ser não negativa (%.6f)", reactance)
	}

	if math.IsNaN(powerFactor) || math.IsInf(powerFactor, 0) {
		return 0, fmt.Errorf("Fator de potência inválido: valor não numérico ou infinito")
	}
	if powerFactor <= 0 || powerFactor > 1 {
		return 0, fmt.Errorf(
			"Fator de potência inválido: deve estar no intervalo (0, 1] (%.2f)",
			powerFactor,
		)
	}

	term :=
		resistance*powerFactor +
			reactance*math.Sqrt(1-powerFactor*powerFactor)

	result :=
		(2 * current * length * term) / 1000

	return result, nil
}
