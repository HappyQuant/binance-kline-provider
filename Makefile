GOCC=go
GOFLAGS=-ldflags '-w -s'

all: rest_server
.PHONY: all

clean: rest_server_clean
.PHONY: clean

rest_server: rest_server_clean
	$(GOCC) build $(GOFLAGS) ./cmd/rest_server
.PHONY: rest_server

rest_server_clean:
	rm -f rest_server
.PHONY: rest_server_clean