package server

type memstore struct {
	store map[uint8]KVData
}

func (m *memstore) Init() {
	m.store = make(map[uint8]KVData)
}

func (m *memstore) put(d *KVData) {
	m.store[d.key] = *d
}

func (m *memstore) get(key uint8) (*KVData, bool) {
	var ret KVData
	e := false
	if val, ok := m.store[key]; ok {
		ret = val
		e = true
	}
	return &ret, e
}
