package generator

import (
	"fmt"
	"github.com/lokstory/digo/model"
	"log"
	"strings"
)

const injectTemplate = `package ${packageName}

import (
	digoContract "${digoContractPath}"
)

func InjectByDigo(manager digoContract.IManager) {
${injectDependency}
}
`

// // Unsafe cast
//	//helloService = *(*hello.IHelloService)(manager.SourcePointer("hello"))

// Unsafe cast
//helloService = *(*hello.IHelloService)(manager.SourcePointer("hello"))

// Reflection
// ptr := manager.SourceValue("hello").Elem().Addr()
// reflect.ValueOf(&helloService).Elem().Set(ptr)

func createInjections(cfg *model.Config) {
	// key: dir
	injectMap := map[string]*model.InjectItem{}
	packageCfg := cfg.PackageConfig

	injectFormat := `	manager.InjectByGeneric("%s", &%s)` + "\n"

	for _, target := range cfg.Targets {
		packageItem := packageCfg.ItemMap[target.FilePath]
		dir := packageItem.Path

		var depth int
		if len(dir) > 0 {
			depth = strings.Count(dir, `/`) + 1
		}
		pathPrefix := pathPrefix(depth)

		item, ok := injectMap[dir]
		if !ok {
			item = &model.InjectItem{
				PackageName:      target.PackageName,
				Dir:              packageItem.Path,
				DigoContractPath: fmt.Sprintf("%s%s", pathPrefix, managerContractDirPath),
			}
			injectMap[dir] = item
		}

		item.InjectDependency += fmt.Sprintf(injectFormat, target.Key, target.VariableName)
	}

	for _, value := range injectMap {
		// remove last empty line
		if len(value.InjectDependency) > 0 {
			value.InjectDependency = strings.TrimSuffix(value.InjectDependency, "\n")
		}

		replacer := strings.NewReplacer(
			`${packageName}`, value.PackageName,
			`${digoContractPath}`, value.DigoContractPath,
			`${injectDependency}`, value.InjectDependency,
		)

		result := replacer.Replace(injectTemplate)

		// generate digo.go
		digoPath := fmt.Sprintf("%s/%s/digo.go", cfg.RootPath, value.Dir)
		if err := saveFile(digoPath, result); err != nil {
			log.Panic(err)
		}
	}
}
