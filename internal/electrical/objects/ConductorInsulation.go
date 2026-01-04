package objects

type ConductorInsulation int

const (
	ConductorInsulationPVC ConductorInsulation = iota
	ConductorInsulationLSHF

	ConductorInsulationXLPE
	ConductorInsulationTRXLPE
	ConductorInsulationEPR
	ConductorInsulationHEPR

	ConductorInsulationEPR105
)
