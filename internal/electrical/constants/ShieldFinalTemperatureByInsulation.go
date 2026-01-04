package constants

import "github.com/lcafe/calculadora_de_cabos/internal/electrical/objects"

var ShieldFinalTemperatureByInsulation = map[objects.ShieldInsulation]float64{
	objects.ShieldInsulationSE1A_SE1B_SHF2: 220.0,
	objects.ShieldInsulationST3:            150.0,
	objects.ShieldInsulationSHF1_ST7:       180.0,
	objects.ShieldInsulationST1_ST2:        200.0,
}
