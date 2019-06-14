package generator

const managerTemplate = `
package digo

import (
${packages}
	"./contract"
	"reflect"
	"sync"
	"unsafe"
)

// Dependency Manager

type manager struct {
	sync.RWMutex
	contract.IManager
	sourceMap map[string]unsafe.Pointer
}

var Manager = &manager{
	sourceMap: map[string]unsafe.Pointer{},
}

func (m *manager) setSource(key string, value interface{}) {
	m.Lock()
	m.sourceMap[key] = unsafe.Pointer(reflect.ValueOf(value).Pointer())
	m.Unlock()
}

func (m *manager) Init() {
	m.initSource()
	m.initDependency()
}
 
func (m *manager) initSource() {
${initSource}
}

func (m *manager) initDependency() {
${initDependency}
}

func (m *manager) GetPointer(key string) unsafe.Pointer {
	m.RLock()
	defer m.RUnlock()
	return m.sourceMap[key]
}
`
