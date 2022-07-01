.PHONY:
.SILENT:

build:
	go build -o ./.bin/online-cinema cmd/online-cinema/main.go
run: build
	sudo ./.bin/online-cinema