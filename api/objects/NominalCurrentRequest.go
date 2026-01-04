package objects

type NominalCurrentRequest struct {
	Power        float64 `json:"power"`
	Voltage      float64 `json:"voltage"`
	PowerFactor  float64 `json:"power_factor"`
	DemandFactor float64 `json:"demand_factor"`
	Efficiency   float64 `json:"efficiency"`
}
