package low_voltage

import (
	"fmt"

	"github.com/lcafe/calculadora_de_cabos/internal/electrical/objects"
	"github.com/lcafe/calculadora_de_cabos/internal/electrical/objects/low_voltage"
)

func NormativeThermalConstants(
	material objects.ConductorMaterial,
	insulation objects.ConductorInsulation,
	sectionMM2 float64,
) (objects.ThermalConstants, error) {

	key := low_voltage.KTableKey{
		Material:       material,
		Insulation:     insulation,
		SectionOver300: sectionMM2 > 300,
	}

	k, ok := kTable[key]
	if !ok {
		return objects.ThermalConstants{}, fmt.Errorf(
			"combinação normativa não suportada para BT (material=%v, isolacao=%v, secao=%.1f)",
			material, insulation, sectionMM2,
		)
	}

	var beta float64
	switch material {
	case objects.ConductorMaterialCopper:
		beta = 234.5
	case objects.ConductorMaterialAluminum:
		beta = 228.0
	default:
		return objects.ThermalConstants{}, fmt.Errorf("material de condutor inválido")
	}

	return objects.ThermalConstants{
		KFactor: k,
		Beta:    beta,
	}, nil
}
