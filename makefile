.SILENT:test
.SILENT:run
.SILENT:build-run
.SILENT:build
.SILENT:js-build
.SILENT:js-run
.SILENT:js-build-run

test:
	go test ./... -coverprofile cover.out && go tool cover -func cover.out
run:
	go run main.go
build-run:
	go build -o service_app && ./service_app
build:
	go build -o service_app
js-build:
	GOOS=js GOARCH=wasm go build -o dist/main.wasm dist/wasm_app.go
js-run:
	http-server -p 8432 dist
js-build-run:
	GOOS=js GOARCH=wasm go build -o dist/main.wasm dist/wasm_app.go
	http-server -p 8432 dist
