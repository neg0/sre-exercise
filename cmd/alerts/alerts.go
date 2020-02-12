package main

import (
	"log"
	"os"

	"uswitch/internal/services"
	"uswitch/internal/services/restful"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("please enter the path to a single JSON file. e.g. ./path/to/service.json")
	}

	_, err := services.NewServices(args[0])
	if err != nil {
		log.Fatal(err)
	}

	err = restful.New()
	if err != nil {
		log.Fatal(err)
	}
}
