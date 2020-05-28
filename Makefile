all: say_hello generate

say_hello:
		@echo "Start build..."

build:
		@echo "Building exec"
		go build main.go

clean:
		@echo "Cleaning up..."