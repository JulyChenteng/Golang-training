package safemap

import "sync"

type SafeMap struct {
	sync.RWMutex
	mp map[interface{}]interface{}
}

func Initial() *SafeMap {
	newMap := new(SafeMap)
	newMap.mp = make(map[interface{}]interface{})
	return newMap
}

func (smp *SafeMap) Read(key interface{}) interface{} {
	smp.RLock()
	value := smp.mp[key]
	smp.RUnlock()

	return value
}

func (smp *SafeMap) Write(key, value interface{}) {
	smp.Lock()
	smp.mp[key] = value
	smp.Unlock()
}

func (smp *SafeMap) Delete(key interface{}) {
	smp.Lock()
	delete(smp.mp, key)
	smp.Unlock()
}

func (smp *SafeMap) Range(f func(key, value interface{}) bool) {
	smp.RLock()
	for k, v := range smp.mp {
		if !f(k, v) {
			break
		}
	}
	smp.RUnlock()
}
