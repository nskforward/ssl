compile:
	# GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags=jsoniter -ldflags="-w -s" -o ./build/app_amd64 ./cmd/api/*
	go build -o ./build/sslgen ./cmd/*

run: compile
	./build/sslgen ca