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
	"log"
	"os"
	"path/filepath"
)

const (
	digoPath = "digo/"
	managerContractPath = digoPath + "contract/manager.go"
)

func Create(cfg *model.Config) {
	createManager(cfg)
}

func deleteIfExists(dir string) error {
	_, err := os.Stat(dir)

	if err == nil {
		return os.Remove(dir)
	}

	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func saveFile(codePath string, codes string) error {
	dir := filepath.Dir(codePath)
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := os.MkdirAll(dir, os.FileMode(0644)); err != nil {
			log.Println(err)
		}
	}

	f, err := os.Create(codePath)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(codes); err != nil {
		return err
	}

	return nil
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
