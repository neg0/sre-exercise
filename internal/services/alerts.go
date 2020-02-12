package services

import (
	"fmt"

	"uswitch/pkg/services"
)

type Alert struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

func NewAlerts(services *services.Services) []Alert {
	var alerts []Alert
	for _, service := range *services {
		for _, objective := range service.Objectives {
			alerts = append(alerts, Alert{
				Name: objective.Name,
				Expression: fmt.Sprintf(
					"%s_%s < %0.2f",
					service.Name,
					objective.Name,
					objective.Target,
				),
			})
		}
	}
	return alerts
}

func NewAlertsByServiceName(services *services.Services, serviceName string) []Alert {
	var alerts []Alert
	for _, service := range *services {
		for _, objective := range service.Objectives {
			if serviceName == service.Name {
				alerts = append(alerts, Alert{
					Name: objective.Name,
					Expression: fmt.Sprintf(
						"%s_%s < %0.2f",
						service.Name,
						objective.Name,
						objective.Target,
					),
				})
			}
		}
	}
	return alerts
}
