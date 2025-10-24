build:
	@go build -o ./bin/bookings ./cmd/web/*.go

run: build
	@./bin/bookings

clean:
	@rm -rf ./bin

test:
	@go test ./...
