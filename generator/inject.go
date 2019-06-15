package generator

import (
	"../model"
	"encoding/json"
	"fmt"
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

type InjectItem struct {
	PackageName string
	DigoContractPath string
	InjectDependency string
}

func createInjections(cfg *model.Config) {
	// key: dir
	injectMap := map[string]*InjectItem{}
	packageCfg := cfg.PackageConfig

	fmt.Println("target length:", len(cfg.Targets))
	for _, target := range cfg.Targets {
		packageItem := packageCfg.ItemMap[target.FilePath]
		dir := packageItem.Path
		fmt.Println("package dir:", packageItem.Path, "alias:", packageItem.Alias)

		var depth int
		if len(dir) > 0 {
			depth = strings.Count(dir, `/`)+2
		}
		pathPrefix := pathPrefix(depth)

		fmt.Println("target path:", target.FilePath, "depth:", depth, "path prefix:", pathPrefix)

		item, ok := injectMap[dir]
		if !ok {
			item = &InjectItem{
				PackageName: target.PackageName,
				DigoContractPath: fmt.Sprintf("%s%s", pathPrefix, managerContractDirPath),
			}
			injectMap[dir] = item
		}
	}

	b, _ := json.Marshal(injectMap)
	fmt.Println("inject map:", string(b))

	// key: filePath
	//templateMap := map[string]string{}
	//packageCfg := cfg.PackageConfig
	//
	//for key, source := range cfg.SourceMap {
	//	packageItem := packageCfg.ItemMap[source.FilePath]
	//
	//	var callPrefix string
	//	alias := packageItem.Alias
	//	if len(alias) > 0 {
	//		callPrefix = alias + "."
	//	}
	//
	//	initSource += fmt.Sprintf(sourceFormat, key, callPrefix, source.VariableName)
	//}
	//
	//for _, target := range cfg.Targets {
	//	packageItem := packageCfg.ItemMap[target.FilePath]
	//	initDependency += fmt.Sprintf(depFormat, packageItem.Alias)
	//
	//	replacer := strings.NewReplacer(
	//		`${packageName}`, target.PackageName,
	//		`${initSource}`, initSource,
	//		`${initDependency}`, initDependency,
	//	)
	//
	//	result := replacer.Replace(template)
	//	fmt.Println("result:", result)
	//
	//	fmt.Println("importPackage:", importPackage)
	//	fmt.Println("initSource:", initSource)
	//	fmt.Println("initDependency:", initDependency)
	//}
	//
	//// generate manager
	//managerPath := fmt.Sprintf("%s/%s/manager.go", cfg.RootPath, ManagerPath)
	//if err := saveFile(managerPath, result); err != nil {
	//	log.Panic(err)
	//}
	//
	//fmt.Println("Generate file successfully")
}
