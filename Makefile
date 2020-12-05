.PHONY: all

all: server client client2

server: 
	$(MAKE) -C cmd/kvs-server

client:
	$(MAKE) -C cmd/kvs-client

client2:
	$(MAKE) -C cmd/kvs-client2

cleanc:
	$(MAKE) -C cmd/kvs-client clean

clean:
	$(MAKE) -C cmd/kvs-server  clean
	$(MAKE) -C cmd/kvs-client clean
	$(MAKE) -C cmd/kvs-client2 clean
