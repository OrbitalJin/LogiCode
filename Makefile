build:
	@go build -o ../bin/LGCC ./cmd/logicode/main.go

run:
	@go run ./cmd/logicode/main.go $(file)
