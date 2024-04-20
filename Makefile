.PHONY: build run clean

# Target binary name
BINARY_NAME=recruitment

# Source files directory
SRC=./cmd/main/main.go

# Default target
build:
	go build -o $(BINARY_NAME) $(SRC)

run:
	./$(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)
