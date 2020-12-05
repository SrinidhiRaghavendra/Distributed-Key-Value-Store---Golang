# ds-assignment3

## build

./configure.remote
make

## run cluster
./startcluster

## start node
./bin/kvs-server config.txt [nodeid [0-3]]] &

## start client
./bin/kvs-client [ip:port]

## distribution of tasks

###Rohit:
-build system
-hint manager
-WAL
-intra cluster communication
-memory store
-client

###Srinidhi:
-quorum manager
-co-ordinator
-thrift file
-replica
-recovery

## completion status
complete
