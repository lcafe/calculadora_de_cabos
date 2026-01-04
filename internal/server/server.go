package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lcafe/calculadora_de_cabos/api"
)

type ServerPort struct {
	Port int `json:"PORT"`
}

func LoadServerPort(path string) (ServerPort, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return ServerPort{}, fmt.Errorf("erro abrindo arquivo %s: %w", path, err)
	}

	var payload ServerPort
	if err := json.Unmarshal(content, &payload); err != nil {
		return ServerPort{}, fmt.Errorf("erro decodificando JSON: %w", err)
	}

	return payload, nil
}

func StartServer() {
	serverPort, serverPortError := LoadServerPort("./internal/server/server.json")
	if serverPortError != nil {
		log.Fatal("Erro no serviço de leitura do server.json: ", serverPortError)
	}

	api.StartRoutes()

	serverAddress := fmt.Sprintf(":%d", serverPort.Port)
	log.Println("Servidor iniciado em", serverAddress)

	if httpServerError := http.ListenAndServe(serverAddress, nil); httpServerError != nil {
		log.Fatal(httpServerError)
	}
}
