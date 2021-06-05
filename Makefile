enemies = 4
build:
	go build -o bin/main main.go

run:
	go run main.go

clean:
	rm ./bin/main

pacman: build
	go run main.go -enemies=$(enemies)