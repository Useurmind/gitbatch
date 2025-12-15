.PHONY: build
build:
	go build .

.PHONY: install
install: build
	sudo cp gitbatch /usr/local/bin