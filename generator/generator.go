package generator

//import (
//	"flag"
//	"log"
//	"os"
//	"path/filepath"
//	"strings"
//)

import (
	"../model"
	"fmt"
	"strings"
)

const (
	ContractPath = "dependency/contract"
	ManagerPath = "dependency"
)

func Create(cfg *model.Config) {
 // packages, initSource, initDependency
	createManager(cfg)
}

func createManager(cfg *model.Config) {
	var packages, initSource, initDependency string
	template := managerTemplate

	sourceFormat := `	m.setSource("%s", &%s)` + "\n"

	for key, item := range cfg.SourceMap {
		initSource += fmt.Sprintf(sourceFormat, key, item.VariableName)
	}

	if len(initSource) > 0 {
		initSource = strings.TrimSuffix(initSource, "\n")
	}

	result := strings.ReplaceAll(template, `${initSource}`, initSource)

	fmt.Println("result:", result)

	fmt.Println("packages:", packages)
	fmt.Println("initSource:", initSource)
	fmt.Println("initDependency:", initDependency)
}

func main() {
	//ymlPath := flag.String("yml", YMLPath, "yaml file path")
	//outputPath := flag.String("output", OutputPath, "output go file path")
	//packageName := flag.String("package", PackageName, "package name")
	//flag.Parse()
	//
	//cfg, err := load(*ymlPath)
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//if len(cfg.Codes) == 0 {
	//	log.Panic("Codes must be set")
	//}
	//
	//codeSet := map[int]struct{}{}
	//keySet := map[string]struct{}{}
	//var successCodeItems string
	//var failCodeItems string
	//var messageItems string
	//
	//for _, item := range cfg.Codes {
	//	if item.Code == nil {
	//		log.Panic("Code must be set")
	//	}
	//
	//	if len(item.Key) == 0 {
	//		log.Panic("Key must be set")
	//	}
	//
	//	code := *item.Code
	//	key := item.Key
	//	message := item.Message
	//
	//	if code == 0 {
	//		log.Panic("Invalid code:", code)
	//	}
	//
	//	if _, ok := codeSet[code]; ok {
	//		log.Panic("Duplicate code:", code)
	//	}
	//
	//	if _, ok := keySet[key]; ok {
	//		log.Panic("Duplicate key:", key)
	//	}
	//
	//	codeSet[code] = struct{}{}
	//	keySet[key] = struct{}{}
	//
	//	codeItem := fmt.Sprintf("    %s Code = %d\n", key, code)
	//
	//	if code >= 0 {
	//		successCodeItems += codeItem
	//	} else {
	//		failCodeItems += codeItem
	//	}
	//	messageItems += fmt.Sprintf("    %s: \"%s\",\n", key, message)
	//}
	//
	//// remove last change line symbols
	//if len(successCodeItems) > 0 {
	//	successCodeItems = strings.TrimSuffix(successCodeItems, "\n")
	//}
	//if len(failCodeItems) > 0 {
	//	failCodeItems = strings.TrimSuffix(failCodeItems, "\n")
	//}
	//if len(keySet) > 0 {
	//	messageItems = strings.TrimSuffix(messageItems, "\n")
	//}
	//
	//fileP := *outputPath
	//dir := filepath.Dir(fileP)
	//if _, err := os.Stat(dir); err != nil {
	//	if !os.IsNotExist(err) {
	//		log.Panic(err)
	//	}
	//	if err := os.MkdirAll(dir, os.FileMode(0644)); err != nil {
	//		log.Println(err)
	//	}
	//}
	//
	//f, err := os.Create(fileP)
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//text := fmt.Sprintf(templateFmt, *packageName, successCodeItems, failCodeItems, messageItems)
	//f.WriteString(text)
	//fmt.Println("Generate file successfully")
}
