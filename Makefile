all: build

build:
		go build -trimpath -o bin/akm ./main.go
