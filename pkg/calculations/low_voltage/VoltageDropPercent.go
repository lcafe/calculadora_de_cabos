package pkg

import (
	"fmt"
	"math"
)

func VoltageDropPercent(
	voltageDrop float64,
	voltage float64,
) (float64, error) {

	if math.IsNaN(voltageDrop) || math.IsInf(voltageDrop, 0) {
		return 0, fmt.Errorf("Queda de tensão inválida: valor não numérico ou infinito")
	}
	if voltageDrop < 0 {
		return 0, fmt.Errorf("Queda de tensão inválida: deve ser não negativa (%.2f)", voltageDrop)
	}

	if math.IsNaN(voltage) || math.IsInf(voltage, 0) {
		return 0, fmt.Errorf("Tensão inválida: valor não numérico ou infinito")
	}
	if voltage <= 0 {
		return 0, fmt.Errorf("Tensão inválida: deve ser maior que zero (%.2f)", voltage)
	}

	return (voltageDrop / voltage) * 100, nil
}
