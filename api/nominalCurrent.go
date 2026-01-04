package api

import (
	"encoding/json"
	"net/http"

	"github.com/lcafe/calculadora_de_cabos/api/objects"
	lv "github.com/lcafe/calculadora_de_cabos/pkg/electrical/calculations/low_voltage"
)

func NominalCurrentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req objects.NominalCurrentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	result, err := lv.NominalCurrent(
		req.Power,
		req.Voltage,
		req.PowerFactor,
		req.DemandFactor,
		req.Efficiency,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := objects.NominalCurrentResponse{
		NominalCurrent: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func nominalCurrentRoute() error {
	http.HandleFunc("/api/nominal-current", NominalCurrentHandler)
	return nil
}
