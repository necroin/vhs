all: build/all start

start:
	bin/unix/vhs -plugins=bin/unix/plugins -log-level=debug

.PHONY: build/all
build/all: build/vhs

.PHONY: build/vhs build/plugins
build/vhs:
	go build -o bin/unix/vhs src/vhs/main.go

.PHONY: build/plugins
build/plugins: build/plugins/filesystem

.PHONY: build/plugins/filesystem
build/plugins/filesystem:
	go build -o bin/unix/plugins/filesystem src/plugins/filesystem/main.go
