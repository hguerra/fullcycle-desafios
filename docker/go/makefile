run:
	go run main.go

build:
	mkdir -p bin/
	go build -o bin/app main.go

build_docker:
	docker build -t heitorcarneiro/fullcycle .

run_docker:
	docker run --rm heitorcarneiro/fullcycle
