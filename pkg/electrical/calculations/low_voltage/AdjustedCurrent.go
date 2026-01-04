package pkg

import (
	"fmt"
	"math"
)

func AdjustedCurrent(nominalCurrent float64, temperatureFactor float64, soilResistivityFactor float64, circuitGroupingFactor float64, serviceFactor float64) (float64, error) {

	if math.IsNaN(nominalCurrent) || math.IsInf((nominalCurrent), 0) {
		return 0, fmt.Errorf("Corrente nominal inválida: valor não numérico ou infinito")
	}

	if nominalCurrent <= 0 {
		return 0, fmt.Errorf(
			"Corrente nominal inválida: deve ser maior que zero (valor informado: %.2f)",
			nominalCurrent,
		)
	}

	if math.IsNaN(temperatureFactor) || math.IsInf(temperatureFactor, 0) {
		return 0, fmt.Errorf("Fator de temperatura inválido: valor não numérico ou infinito")
	}

	if temperatureFactor <= 0 {
		return 0, fmt.Errorf(
			"Fator de temperatura inválido: deve ser maior que zero (valor informado: %.2f)",
			temperatureFactor,
		)
	}

	if temperatureFactor > 1 {
		return 0, fmt.Errorf(
			"Fator de temperatura inválido: não pode ser maior que 1 conforme normas técnicas (valor informado: %.2f)",
			temperatureFactor,
		)
	}

	if math.IsNaN(soilResistivityFactor) || math.IsInf(soilResistivityFactor, 0) {
		return 0, fmt.Errorf("Fator de resistividade do solo inválido: valor não numérico ou infinito")
	}

	if soilResistivityFactor <= 0 {
		return 0, fmt.Errorf(
			"Fator de resistividade do solo inválido: deve ser maior que zero (valor informado: %.2f)",
			soilResistivityFactor,
		)
	}

	if math.IsNaN(circuitGroupingFactor) || math.IsInf(circuitGroupingFactor, 0) {
		return 0, fmt.Errorf("Fator de agrupamento de circuitos inválido: valor não numérico ou infinito")
	}

	if circuitGroupingFactor <= 0 {
		return 0, fmt.Errorf(
			"Fator de agrupamento de circuitos inválido: deve ser maior que zero (valor informado: %.2f)",
			circuitGroupingFactor,
		)
	}

	if math.IsNaN(serviceFactor) || math.IsInf(serviceFactor, 0) {
		return 0, fmt.Errorf("Fator de serviço inválido: valor não numérico ou infinito")
	}

	if serviceFactor < 1 {
		return 0, fmt.Errorf(
			"Fator de serviço inválido: deve ser maior ou igual a um (valor informado: %.2f)",
			serviceFactor,
		)
	}

	result := (nominalCurrent * serviceFactor) /
		(temperatureFactor * soilResistivityFactor * circuitGroupingFactor)

	return result, nil
}
