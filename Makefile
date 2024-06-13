cms:
	go build -o bin/go-cms cmd/goCms/main.go

cli:
	go build -o bin/go-cli cmd/goCms/main.go

templ:
	TEMPL_EXPERIMENT=rawgo templ generate

dev: templ
	air --build.cmd "go build -o bin/go-cms cmd/goCms/main.go" --build.bin "./bin/go-cms"

cert:
	openssl req -newkey rsa:4096 \
            -x509 \
            -sha256 \
            -days 3650 \
            -nodes \
            -out cert/tmp.crt \
            -keyout cert/tmp.key

all: cms cli
