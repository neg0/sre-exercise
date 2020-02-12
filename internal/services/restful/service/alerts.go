package service

import (
	"encoding/json"
	"sync"

	internalServices "uswitch/internal/services"
	"uswitch/pkg/services"
)

func AlertsHandler(serviceName string, allService *services.Services) ([]byte, error) {
	var body []byte
	var err error

	if allService == nil {
		return body, err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		alerts := internalServices.NewAlerts(allService)
		if len(serviceName) > 0 {
			alerts = internalServices.NewAlertsByServiceName(
				allService,
				serviceName,
			)
		}
		body, err = json.Marshal(alerts)
		wg.Done()
	}()
	wg.Wait()
	return body, err
}
