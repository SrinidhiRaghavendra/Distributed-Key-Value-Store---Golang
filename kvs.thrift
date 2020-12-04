exception SystemException {
  1: optional string message
}

struct KVData {
  1: string value;
  2: i64 timestamp;
  3: i32 key;
}

struct Node {
  1: i32 id;
  2: string ip;
  3: i32 port;
}

enum ConsistencyLevel {
  ONE,
  QUORUM
}

service Replica {
  string get(1: i32 key, 2: ConsistencyLevel cLevel)
    throws (1: SystemException systemException),
  
  void put(1: i32 key, 2: string value, 3: ConsistencyLevel cLevel)
    throws (1: SystemException systemException),
  
  KVData getDataFromNode(1: i32 key) 
    throws (1: SystemException systemException),

  void putDataInNode(1: KVData data) 
    throws (1: SystemException systemException),

  list<KVData> getHints(1: Node node) 
    throws (1: SystemException systemException),
}

