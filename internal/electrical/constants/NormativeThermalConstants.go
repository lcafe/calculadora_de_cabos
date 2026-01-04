package constants

import "github.com/lcafe/calculadora_de_cabos/internal/electrical/objects"

var (
	ThermalConstantsCopper = objects.ThermalConstants{
		KFactor: 226.0,
		Beta:    234.5,
	}

	ThermalConstantsAluminum = objects.ThermalConstants{
		KFactor: 148.0,
		Beta:    228.0,
	}
)
