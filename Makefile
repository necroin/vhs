all: build/all start

start:
	bin/unix/vhs.exe -plugins=bin/unix/plugins -log-level=debug

.PHONY: build/all
build/all: build/vhs

.PHONY: build/vhs
build/vhs:
	go build -o bin/unix/vhs.exe src/vhs/main.go
