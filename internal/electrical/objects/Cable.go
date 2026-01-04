package objects

type Cable struct {
	NominalSectionMM2 float64
	LengthMeters      float64

	ResistanceOhmKm float64
	ReactanceOhmKm  float64

	ConductorMaterial   ConductorMaterial
	ConductorInsulation ConductorInsulation

	Thermal ThermalConstants
}
