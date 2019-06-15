package generator

//import (
//	"flag"
//	"log"
//	"os"
//	"path/filepath"
//	"strings"
//)

import (
	"github.com/lokstory/digo/model"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	digoDirPath             = "digo"
	managerContractDirPath  = digoDirPath + "/contract"
	managerContractFilePath = managerContractDirPath + "/manager.go"
)

// Start to generate manager and digo file
func Start(cfg *model.Config) {
	createManager(cfg)
	createInjections(cfg)
}

// delete file if exists
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

// get the path to root by depth
// example: hello/world will return ../../
func pathPrefix(depth int) string {
	if depth <= 0 {
		return ""
	}

	return strings.Repeat(`../`, depth)
}

// save go file by string
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
