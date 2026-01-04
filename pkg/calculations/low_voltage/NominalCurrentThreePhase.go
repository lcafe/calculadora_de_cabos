package pkg

import (
	"fmt"
	"math"
)

func NominalCurrentThreePhase(
	power float64,
	voltage float64,
	powerFactor float64,
	demandFactor float64,
) (float64, error) {

	if math.IsNaN(power) || math.IsInf(power, 0) {
		return 0, fmt.Errorf("Potência inválida: valor não numérico ou infinito")
	}
	if power <= 0 {
		return 0, fmt.Errorf(
			"Potência inválida: deve ser maior que zero (valor informado: %.2f)",
			power,
		)
	}

	if math.IsNaN(voltage) || math.IsInf(voltage, 0) {
		return 0, fmt.Errorf("Tensão inválida: valor não numérico ou infinito")
	}
	if voltage <= 0 {
		return 0, fmt.Errorf(
			"Tensão inválida: deve ser maior que zero (valor informado: %.2f)",
			voltage,
		)
	}

	if math.IsNaN(powerFactor) || math.IsInf(powerFactor, 0) {
		return 0, fmt.Errorf("Fator de potência inválido: valor não numérico ou infinito")
	}
	if powerFactor <= 0 {
		return 0, fmt.Errorf(
			"Fator de potência inválido: deve ser maior que zero (valor informado: %.2f)",
			powerFactor,
		)
	}
	if powerFactor > 1 {
		return 0, fmt.Errorf(
			"Fator de potência inválido: deve ser menor ou igual a 1 (valor informado: %.2f)",
			powerFactor,
		)
	}

	if math.IsNaN(demandFactor) || math.IsInf(demandFactor, 0) {
		return 0, fmt.Errorf("Fator de demanda inválido: valor não numérico ou infinito")
	}
	if demandFactor <= 0 {
		return 0, fmt.Errorf(
			"Fator de demanda inválido: deve ser maior que zero (valor informado: %.2f)",
			demandFactor,
		)
	}
	if demandFactor > 1 {
		return 0, fmt.Errorf(
			"Fator de demanda inválido: deve ser menor ou igual a 1 (valor informado: %.2f)",
			demandFactor,
		)
	}

	result := (power * demandFactor) /
		(math.Sqrt(3) * voltage * powerFactor)

	return result, nil
}
