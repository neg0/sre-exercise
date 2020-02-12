package service

import (
	"encoding/json"
	"sync"

	"uswitch/pkg/services"
)

func ObjectivesHandler(serviceName string, allService *services.Services) ([]byte, error) {
	var body []byte
	var err error

	if allService == nil {
		return body, err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		body, err = json.Marshal(allService)

		if len(serviceName) > 0 {
			var allServices services.Services
			for _, service := range *allService {
				if service.Name == serviceName {
					allServices = append(allServices, service)
					body, err = json.Marshal(allServices)
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()
	return body, err
}
