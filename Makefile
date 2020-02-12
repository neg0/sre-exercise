.PHONY: help     	# Generate list of targets with descriptions
help:
	@echo "\n"
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | expand -t20


.PHONY: compile 	# Compiles the main app
compile:
	rm -rf dist
	go build -o dist/alerts cmd/alerts/alerts.go


.PHONY: run    		# runs the main app
run:
	./dist/alerts ./service.json


.PHONY: build  		# Builds Docker images for container registery
build:
	docker build -t uswitch-container.dev/techtest:1.0 .


.PHONY: up    		# runs the container
up:
	docker run --rm -d -p 8080:8080 --name hadi-tech-test uswitch-container.dev/techtest:1.0


.PHONY: down    	# stops the running container
down:
	docker stop hadi-tech-test