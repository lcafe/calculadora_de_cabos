package low_voltage

import (
	"fmt"
	"math"
)

func NominalCurrentPerCable(
	nominalCurrent float64,
	cablePerPhase int,
) (float64, error) {

	if math.IsNaN(nominalCurrent) || math.IsInf(nominalCurrent, 0) {
		return 0, fmt.Errorf("Corrente nominal inválida: valor não numérico ou infinito")
	}

	if cablePerPhase <= 0 {
		return 0, fmt.Errorf("Quantidade de cabos por fase inválida: deve ser maior que zero")
	}

	result := nominalCurrent / float64(cablePerPhase)

	return result, nil
}
