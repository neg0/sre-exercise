# Site Reliability Engineer Exercise
* Golang 1.12+


## Architecture
I have utilised Effective Go and Uber Go Style Guide to produce this solution. App is created using Domain-Driven Design
(Rich Domain Model) with implementation of SOLID principles.


## Build, and Run on Local Host
Please ensure you have a Golang 1.12+ installed on your host machine if not you could use the docker implementation in this project. 
First ensure you are at the root of this repository.
To download the dependencies, please run:

    ~$: go mod download

and then to compile please run:
    
    ~$: make compile

and now you can run the executable at the root of the project:
    
    ~$: make run


## Build, and Run using Docker
I created a `Makefile` commands for convenience of running docker commands.

    ~$: make build
    ~$: make up
    ~$: make down


### API
Either from Local Host or Docker, API should be running on `localhost` port `8080`. e.g. to view the Alerts visit: 

    http://127.0.0.1:8080/alerts/energy


### Tests
Do to time constrains, I haven't added many tests for this solution, but I hope it's enough for purpose of this assessment.


#### Credits
 * [Hadi Tajallaei](mailto:tajallaei@gmail.com)