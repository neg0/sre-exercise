package services

import (
	"io/ioutil"
	"sync"

	"uswitch/pkg/services"
)

var (
	once sync.Once
	err  error

	instance *services.Services
)

func Instance() *services.Services {
	return instance
}

func NewServices(filePath string) (*services.Services, error) {
	once.Do(func() {
		contents, errRead := ioutil.ReadFile(filePath)
		if errRead != nil {
			err = errRead
			return
		}

		servicesInstance, errJSON := services.NewServicesFromJSON(contents)
		if errJSON != nil {
			err = errJSON
			return
		}
		instance = servicesInstance
	})
	return instance, err
}
