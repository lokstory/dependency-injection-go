package generator

import (
	"fmt"
	"github.com/lokstory/digo/model"
	"log"
	"strings"
)

const managerTemplate = `package digo

import (
${importPackage}
	"./contract"
	"reflect"
	"sync"
	"unsafe"
)

// Digo Dependency Manager
type Manager struct {
	sync.RWMutex
	contract.IManager
	sourceValueMap map[string]reflect.Value
	sourcePtrMap map[string]unsafe.Pointer
}

// Singleton
var manager = &Manager{
	sourceValueMap: map[string]reflect.Value{},
	sourcePtrMap: map[string]unsafe.Pointer{},
}

// Start to set source, then inject dependencies
func Start() {
	manager.initSources()
	manager.injectDependencies()
}

// Inject dependencies
func (m *Manager) injectDependencies() {
${initDependency}
}

// Initialize sources
func (m *Manager) initSources() {
${initSource}
}

// set the source by key
func (m *Manager) setSource(key string, source interface{}) {
	value := reflect.ValueOf(source)
	
	m.Lock()
	
	m.sourceValueMap[key] = value
	m.sourcePtrMap[key] = unsafe.Pointer(value.Pointer())
	
	m.Unlock()
}

// SourceValue get the source's pointer by key
func (m *Manager) SourcePointer(key string) unsafe.Pointer {
	m.RLock()
	defer m.RUnlock()
	return m.sourcePtrMap[key]
}

// SourceValue get the source's generic value by key
func (m *Manager) SourceValue(key string) reflect.Value {
	m.RLock()
	defer m.RUnlock()
	return m.sourceValueMap[key]
}

// InjectByValue inject the dependency by source's key and target generic value
func (m *Manager) InjectByGeneric(sourceKey string, value interface{}) {
	source := m.SourceValue(sourceKey).Elem()
	reflect.ValueOf(value).Elem().Set(source)
}

`

// Create dependency manager and contract in the digo package of project's root path
func createManager(cfg *model.Config) {
	var importPackage, initSource, initDependency string
	template := managerTemplate
	packageCfg := cfg.PackageConfig

	packageFormat := `	%s"../%s"` + "\n"
	sourceFormat := `	m.setSource("%s", &%s%s)` + "\n"
	depFormat := `	%s.InjectByDigo(m)` + "\n"

	for key, source := range cfg.SourceMap {
		packageItem := packageCfg.ItemMap[source.FilePath]

		var callPrefix string
		alias := packageItem.Alias
		if len(alias) > 0 {
			callPrefix = alias + "."
		}

		initSource += fmt.Sprintf(sourceFormat, key, callPrefix, source.VariableName)
	}

	for _, target := range cfg.Targets {
		packageItem := packageCfg.ItemMap[target.FilePath]
		initDependency += fmt.Sprintf(depFormat, packageItem.Alias)
	}

	// Sort packages
	for _, dir := range packageCfg.Packages {
		packageItem := packageCfg.DirMap[dir]
		importPackage += fmt.Sprintf(packageFormat, packageItem.Alias+" ", dir)
	}

	// Remove the last empty line if exists
	if len(importPackage) > 0 {
		importPackage = strings.TrimSuffix(importPackage, "\n")
	}
	if len(initSource) > 0 {
		initSource = strings.TrimSuffix(initSource, "\n")
	}
	if len(initDependency) > 0 {
		initDependency = strings.TrimSuffix(initDependency, "\n")
	}

	replacer := strings.NewReplacer(
		`${importPackage}`, importPackage,
		`${initSource}`, initSource,
		`${initDependency}`, initDependency,
	)

	result := replacer.Replace(template)

	// generate contract
	contractPath := fmt.Sprintf("%s/%s", cfg.RootPath, managerContractFilePath)
	if err := saveFile(contractPath, ContractTemplate); err != nil {
		log.Panic(err)
	}

	// generate manager
	managerPath := fmt.Sprintf("%s/%s/manager.go", cfg.RootPath, digoDirPath)
	if err := saveFile(managerPath, result); err != nil {
		log.Panic(err)
	}

	fmt.Println("Generate file successfully")
}
