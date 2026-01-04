package low_voltage

import "github.com/lcafe/calculadora_de_cabos/internal/electrical/objects"

type KTableKey struct {
	Material       objects.ConductorMaterial
	Insulation     objects.ConductorInsulation
	SectionOver300 bool
}
