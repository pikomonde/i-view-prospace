.SILENT:test
.SILENT:run
.SILENT:build-run
.SILENT:build

test:
	go test ./... -coverprofile cover.out && go tool cover -func cover.out
run:
	go run main.go
build-run:
	go build -o service_app && ./service_app
build:
	go build -o service_app
