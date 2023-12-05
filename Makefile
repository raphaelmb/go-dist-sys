LOG_BINARY=log
REGISTRY_BINARY=registry
GRADING_BINARY=grading

build_log:
	@echo "Building log binary..."
	@env GOOS=linux CGO_ENABLED=0 go build -o bin/${LOG_BINARY} ./cmd/logservice
	@echo "Done!"

build_registry:
	@echo "Building registry binary..."
	@env GOOS=linux CGO_ENABLED=0 go build -o bin/${REGISTRY_BINARY} ./cmd/registryservice
	@echo "Done!"

build_grading:
	@echo "Building grading binary..."
	@env GOOS=linux CGO_ENABLED=0 go build -o bin/${GRADING_BINARY} ./cmd/gradingservice
	@echo "Done!"

.PHONY: build_log build_registry build_grading