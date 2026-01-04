package objects

func (m ConductorMaterial) String() string {
	switch m {
	case ConductorMaterialCopper:
		return "cobre"
	case ConductorMaterialAluminum:
		return "aluminio"
	default:
		return "desconhecido"
	}
}

func (i ConductorInsulation) String() string {
	switch i {

	case ConductorInsulationPVC:
		return "pvc"
	case ConductorInsulationLSHF:
		return "lshf"

	case ConductorInsulationXLPE:
		return "xlpe"
	case ConductorInsulationTRXLPE:
		return "tr-xlpe"
	case ConductorInsulationEPR:
		return "epr"
	case ConductorInsulationHEPR:
		return "hepr"

	case ConductorInsulationEPR105:
		return "epr-105"

	default:
		return "desconhecido"
	}
}
