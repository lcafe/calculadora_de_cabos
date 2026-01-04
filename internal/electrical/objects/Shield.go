package objects

type Shield struct {
	InsulationType      ShieldInsulation
	InitialTemperatureC float64
	FinalTemperatureC   float64

	Thermal ThermalConstants
}
