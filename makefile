test:
	go test ./... -coverprofile cover.out && go tool cover -func cover.out
run:
	go run main.go
build-run:
	go build -o service_app && ./service_app
