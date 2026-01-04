package pkg

import (
	"fmt"
	"math"
)

func CheckOverloadProtection(
	designCurrent float64,
	breakerRatedCurrent float64,
	breakerTripCurrent float64,
	correctedCableCurrent float64,
	cableThermalLimit float64,
) (bool, error) {

	if math.IsNaN(designCurrent) || math.IsInf(designCurrent, 0) {
		return false, fmt.Errorf("corrente de projeto inválida: valor não numérico ou infinito")
	}
	if designCurrent < 0 {
		return false, fmt.Errorf(
			"corrente de projeto inválida: valor negativo (%.2f)",
			designCurrent,
		)
	}

	if math.IsNaN(breakerRatedCurrent) || math.IsInf(breakerRatedCurrent, 0) {
		return false, fmt.Errorf("corrente nominal do disjuntor inválida: valor não numérico ou infinito")
	}
	if breakerRatedCurrent < 0 {
		return false, fmt.Errorf(
			"corrente nominal do disjuntor inválida: valor negativo (%.2f)",
			breakerRatedCurrent,
		)
	}

	if math.IsNaN(breakerTripCurrent) || math.IsInf(breakerTripCurrent, 0) {
		return false, fmt.Errorf("corrente de atuação do disjuntor inválida: valor não numérico ou infinito")
	}
	if breakerTripCurrent < 0 {
		return false, fmt.Errorf(
			"corrente de atuação do disjuntor inválida: valor negativo (%.2f)",
			breakerTripCurrent,
		)
	}

	if math.IsNaN(correctedCableCurrent) || math.IsInf(correctedCableCurrent, 0) {
		return false, fmt.Errorf("corrente admissível do cabo inválida: valor não numérico ou infinito")
	}
	if correctedCableCurrent < 0 {
		return false, fmt.Errorf(
			"corrente admissível do cabo inválida: valor negativo (%.2f)",
			correctedCableCurrent,
		)
	}

	if math.IsNaN(cableThermalLimit) || math.IsInf(cableThermalLimit, 0) {
		return false, fmt.Errorf("limite térmico do cabo inválido: valor não numérico ou infinito")
	}
	if cableThermalLimit < 0 {
		return false, fmt.Errorf(
			"limite térmico do cabo inválido: valor negativo (%.2f)",
			cableThermalLimit,
		)
	}

	var result bool

	conditionA :=
		designCurrent <= breakerRatedCurrent &&
			breakerRatedCurrent <= correctedCableCurrent

	conditionB :=
		breakerTripCurrent <= cableThermalLimit

	if conditionA && conditionB {
		result = true
	} else {
		result = false
	}

	return result, nil
}
