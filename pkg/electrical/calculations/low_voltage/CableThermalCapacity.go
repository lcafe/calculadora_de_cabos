package low_voltage

import (
	"fmt"
	"math"
)

func CableThermalCapacity(
	kFactor float64,
	section float64,
	cablesPerPhase int,
) (float64, error) {

	if math.IsNaN(kFactor) || math.IsInf(kFactor, 0) {
		return 0, fmt.Errorf("Constante K inválida: valor não numérico ou infinito")
	}
	if kFactor < 0 {
		return 0, fmt.Errorf("Constante K inválida: deve ser maior ou igual a zero (%.2f)", kFactor)
	}

	if math.IsNaN(section) || math.IsInf(section, 0) {
		return 0, fmt.Errorf("Seção inválida: valor não numérico ou infinito")
	}
	if section < 0 {
		return 0, fmt.Errorf("Seção inválida: deve ser maior ou igual a zero (%.2f)", section)
	}

	if cablesPerPhase <= 0 {
		return 0, fmt.Errorf("Quantidade de cabos por fase inválida: deve ser maior que zero")
	}

	equivalentSection := section * float64(cablesPerPhase)

	result := math.Pow(kFactor, 2) * math.Pow(equivalentSection, 2)
	return result, nil
}
