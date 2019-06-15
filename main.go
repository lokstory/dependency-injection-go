package main

import (
	"flag"
	"github.com/lokstory/digo/generator"
	"github.com/lokstory/digo/model"
	"log"
	"os"
	"path/filepath"
)

// TODO CLI
// Generate digo package and inject files
func build() {

}

// TODO CLI
// Delete digo package and all of inject files
func delete() {

}

// Start
func main() {
	var rootPath string
	flag.StringVar(&rootPath, "path", "", "root path of project")

	flag.Parse()

	if len(rootPath) == 0 {
		log.Panicln("project path must be set")
	}

	if _, err := os.Stat(rootPath); err != nil {
		log.Panicln(err)
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

	generator.Start(cfg)
}
