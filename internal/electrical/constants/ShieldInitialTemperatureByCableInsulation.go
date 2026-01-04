package constants

import "github.com/lcafe/calculadora_de_cabos/internal/electrical/objects"

var ShieldInitialTemperatureByCableInsulation = map[objects.ConductorInsulation]float64{
	objects.ConductorInsulationXLPE:   85.0,
	objects.ConductorInsulationTRXLPE: 85.0,
	objects.ConductorInsulationEPR:    85.0,
	objects.ConductorInsulationHEPR:   85.0,

	objects.ConductorInsulationEPR105: 100.0,
}
