package services_test

import (
	"fmt"

	. "uswitch/pkg/services"
)

func ExampleNewServicesFromJSON() {
	data := `[
	  {
		"name": "energy",
		"objectives": [
		  {
			"name": "availability",
			"target": 0.99
		  },
		  {
			"name": "latency",
			"target": 0.95
		  }
		]
	  },
	  {
		"name": "mobiles",
		"objectives": [
		  {
			"name": "availability",
			"target": 0.95
		  },
		  {
			"name": "latency",
			"target": 0.95
		  }
		]
	  },
	  {
		"name": "deals",
		"objectives": [
		  {
			"name": "correctness",
			"target": 0.95
		  },
		  {
			"name": "freshness",
			"target": 0.95
		  }
		]
	  }
	]
	`

	services, _ := NewServicesFromJSON([]byte(data))

	for _, service := range *services {
		fmt.Println(service.Name)
	}

	// Output:
	// energy
	// mobiles
	// deals
}
