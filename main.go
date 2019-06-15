package main

import (
	"./generator"
	"./model"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

// Generate
func build() {

}

// Delete digo package and all of digo.go
func delete() {

}

func main() {
	var rootPath string
	flag.StringVar(&rootPath, "path", ".", "root path of project")

	flag.Parse()

	if len(rootPath) == 0 {
		log.Panic("path must be set")
	}

	if p, err := filepath.Abs(rootPath); err != nil {
		log.Panic(err)
	} else {
		rootPath = p
	}

	cfg, err := model.CreateConfig(rootPath)
	if err != nil {
		log.Panic(err)
	}

	b, _ := json.Marshal(cfg)
	fmt.Println("config:", string(b))

	generator.Start(cfg)
}
