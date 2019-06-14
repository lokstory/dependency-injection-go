package main

import (
	"./generator"
	"./model"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func regexVer() {
	//testReg := regexp.MustCompile(`//\s*@Source\((\w+)\)[\s\t\r\n]+var\s+(\w+)\s+(\w+)\s+=\s+&*([\w]+)\s*`)
	//log.Println("text length:", len(l), err)
	//log.Println("test:", testReg.FindAllStringSubmatch(string(l), -1))
}

var (
	rootPath string
)

func parseItems() (sourceMap map[string]*model.AnnotationItem, targets []*model.AnnotationItem, retErr error) {
	sourceMap = map[string]*model.AnnotationItem{}

	annotationReg := regexp.MustCompile(`//\s*@(\w+)\(*(\w*)\)*`)
	variableReg := regexp.MustCompile(`var\s+(\w+)\s+([.\w]+)[\s=&]*([\w]*)`)

	retErr = filepath.Walk(rootPath, func(p string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(p, ".go") || strings.HasPrefix(p, "_test.go") {
			return nil
		}

		fmt.Println(p)

		f, err := os.Open(p)
		if err != nil {
			return err
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		var item *model.AnnotationItem
		var reset = func() {
			item = nil
		}

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}

			text := strings.TrimSpace(string(line))

			// Skip empty line
			if len(text) == 0 {
				reset()
				continue
			}

			if strings.HasPrefix(text, "// @") {
				matches := annotationReg.FindStringSubmatch(text)
				if len(matches) < 2 {
					continue
				}

				annotation := matches[1]

				switch annotation {
				case model.AnnotationSource, model.AnnotationInject:
					var injectType string
					var key string

					if len(matches) >= 3 && len(matches[2]) > 0 {
						injectType = model.InjectByKey
						key = matches[2]
					} else {
						injectType = model.InjectByType
					}

					item = &model.AnnotationItem{
						FilePath:       p,
						InjectType:     injectType,
						AnnotationType: annotation,
						Key:            key,
					}
				default:
					//reset()
				}

				continue
			}

			if item == nil || strings.HasPrefix(text, "//") {
				continue
			}

			matches := variableReg.FindStringSubmatch(text)
			if len(matches) < 2 {
				log.Panic("invalid variable:", text)
			}

			var typeName string

			// remove package
			if len(matches) >= 3 {
				typeName = matches[2]
				if id := strings.Index(typeName, "."); id >= 0 {
					typeName = typeName[id+1:]
				}
			}

			item.VariableName = matches[1]
			item.TypeName = typeName
			if item.InjectType == model.InjectByType {
				item.Key = typeName
			}

			switch item.AnnotationType {
			case model.AnnotationSource:
				key := item.Key
				if _, ok := sourceMap[key]; ok {
					log.Panicf(
						"duplicate source, key: %s, inject type: %s",
						key,
						item.InjectType,
					)
				}

				sourceMap[key] = item
			case model.AnnotationInject:
				targets = append(targets, item)
			}

			reset()
			continue
		}

		return nil
	})

	return
}

func main() {
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

	sourceMap, targets, err := parseItems()
	if err != nil {
		log.Panic(err)
	}

	b, _ := json.Marshal(sourceMap)
	fmt.Println("sources:", string(b))

	b, _ = json.Marshal(targets)
	fmt.Println("targets:", string(b))

	cfg := &model.Config{
		RootPath:  rootPath,
		SourceMap: sourceMap,
		Targets:   targets,
	}

	generator.Create(cfg)
}
