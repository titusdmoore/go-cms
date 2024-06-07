cms:
	go build -o bin/go-cms cmd/goCms/main.go

cli:
	go build -o bin/go-cli cmd/goCms/main.go

all: cms cli
