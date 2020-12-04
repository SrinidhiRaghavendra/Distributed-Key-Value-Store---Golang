.PHONY: all

all: server client

server: 
	$(MAKE) -C cmd/kvs-server

client:
	$(MAKE) -C cmd/kvs-client

cleanc:
	$(MAKE) -C cmd/kvs-client clean

clean:
	$(MAKE) -C cmd/kvs-server  clean
	$(MAKE) -C cmd/kvs-client clean
