package dependency

import (
	//"../example/hello"
	//"../example/world"
	"./contract"
	"reflect"
	"sync"
	"unsafe"
)

// Dependency Manager

type manager struct {
	sync.RWMutex
	contract.IManager
	sourceValueMap map[string]reflect.Value
	sourcePtrMap map[string]unsafe.Pointer
}

var Manager = &manager{
	sourceValueMap: map[string]reflect.Value{},
	sourcePtrMap: map[string]unsafe.Pointer{},
}

func (m *manager) setSource(key string, source interface{}) {
	value := reflect.ValueOf(source)

	m.Lock()

	m.sourceValueMap[key] = value
	m.sourcePtrMap[key] = unsafe.Pointer(value.Pointer())

	m.Unlock()
}

func (m *manager) Init() {
	m.initSource()
	m.initDependency()
}

func (m *manager) initSource() {
}

func (m *manager) initDependency() {
}

func (m *manager) SourcePointer(key string) unsafe.Pointer {
	m.RLock()
	defer m.RUnlock()
	return m.sourcePtrMap[key]
}

func (m *manager) SourceValue(key string) reflect.Value {
	m.RLock()
	defer m.RUnlock()
	return m.sourceValueMap[key]
}
