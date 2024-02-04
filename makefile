run:
	@echo "Building trai..."
	@mkdir -p .build
	@go build -o .build/trai_local main.go
	@sudo cp ./.build/trai_local ~/../../usr/local/bin/
	


