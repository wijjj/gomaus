# Go params
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean

BINARY_NAME=gomaus
MAIN_FILE=./cmd/main.go

# Dependency installation for Robotgo (Debian)
install-deps:
	@dpkg -l libx11-dev libxtst-dev libx11-xcb-dev libxkbcommon-dev libxkbcommon-x11-dev >/dev/null || ( \
		echo "Installing missing dependencies..." && \
		sudo apt-get update && \
		sudo apt-get install -y libx11-dev libxtst-dev libx11-xcb-dev libxkbcommon-dev libxkbcommon-x11-dev \
	)

# Build the project
build: install-deps
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

# Run the project
run: build
	./$(BINARY_NAME)

# Clean up the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Default rule (install deps, build and run)
all: run

.PHONY: build run clean all
