package dawn

import (
	"sync"
)

type ConnMap struct {
	sync.RWMutex
	m map[int64]*ServerConn
}

// NewConnMap returns a new ConnMap.
func NewConnMap() *ConnMap {
	return &ConnMap{
		m: make(map[int64]*ServerConn),
	}
}

// Clear clears all elements in map.
func (cm *ConnMap) Clear() {
	cm.Lock()
	cm.m = make(map[int64]*ServerConn)
	cm.Unlock()
}

// Get gets a server connection with specified net ID.
func (cm *ConnMap) Get(id int64) (*ServerConn, bool) {
	cm.RLock()
	sc, ok := cm.m[id]
	cm.RUnlock()
	return sc, ok
}

// Put puts a server connection with specified net ID in map.
func (cm *ConnMap) Put(id int64, sc *ServerConn) {
	cm.Lock()
	cm.m[id] = sc
	cm.Unlock()
}

// Remove removes a server connection with specified net ID.
func (cm *ConnMap) Remove(id int64) {
	cm.Lock()
	delete(cm.m, id)
	cm.Unlock()
}

// Size returns map size.
func (cm *ConnMap) Size() int {
	cm.RLock()
	size := len(cm.m)
	cm.RUnlock()
	return size
}

// IsEmpty tells whether ConnMap is empty.
func (cm *ConnMap) IsEmpty() bool {
	return cm.Size() <= 0
}