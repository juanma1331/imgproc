.PHONY: build
build:
	@go build -o ./bin/imgproc.exe ./main.go
	@echo "Build complete!"

.PHONY: clone
clone: 
	@go run ./cmd/clone-images/main.go
