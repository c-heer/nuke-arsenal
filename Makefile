.PHONY: dev build install clean frontend bindings

# Development
dev:
	cd frontend && npm install && npm run build
	wails3 dev

# Build production binary
build: frontend bindings
	go build -o arsenal ./cmd/arsenal

# Install to /usr/local/bin
install: build
	cp arsenal /usr/local/bin/

# Build frontend only
frontend:
	cd frontend && npm install && npm run build

# Generate Wails bindings
bindings:
	wails3 generate bindings -d frontend/bindings -clean=true ./cmd/arsenal/...

# Clean build artifacts
clean:
	rm -rf arsenal frontend/dist frontend/node_modules
