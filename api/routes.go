package api

import "fmt"

func StartRoutes() error {
	if helloRouteError := helloRoute(); helloRouteError != nil {
		return fmt.Errorf("falha ao registrar rota /hello: %w", helloRouteError)
	}
	return nil
}
