test:
	go test ./... -coverprofile cover.out && go tool cover -func cover.out
run:
	go run main.go