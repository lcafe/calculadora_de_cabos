package low_voltage

import (
	"fmt"
	"math"
)

func MotorVoltageDrop(
	current float64,
	length float64,
	resistance float64,
	cablesPerPhase int,
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

	if cablesPerPhase <= 0 {
		return 0, fmt.Errorf("Quantidade de cabos por fase inválida: deve ser maior que zero")
	}

	result :=
		(2 * current * length * (resistance * 0.3)) /
			(float64(cablesPerPhase) * 1000)

	return result, nil
}
