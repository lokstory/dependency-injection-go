package model

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Config struct {
	RootPath      string
	SourceMap     map[string]*AnnotationItem
	Targets       []*AnnotationItem
	PackageConfig *PackageConfig
}

func CreateConfig(rootPath string) (cfg *Config, retErr error) {
	sourceMap := map[string]*AnnotationItem{}
	var targets []*AnnotationItem

	annotationReg := regexp.MustCompile(`//\s*@(\w+)\(*(\w*)\)*`)
	variableReg := regexp.MustCompile(`var\s+(\w+)\s+([.\w]+)[\s=&]*([\w]*)`)

	retErr = filepath.Walk(rootPath, func(p string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(p, ".go") || strings.HasPrefix(p, "_test.go") {
			return nil
		}

		f, err := os.Open(p)
		if err != nil {
			return err
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		var packageName string
		var item *AnnotationItem
		var reset = func() {
			item = nil
		}

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}

			text := strings.TrimSpace(string(line))

			// Skip and disallow empty lines
			if len(text) == 0 {
				reset()
				continue
			}

			if len(packageName) == 0 &&
				strings.HasPrefix(text, "package ") {
				packageName = strings.TrimPrefix(text, "package ")
				continue
			}

			if strings.HasPrefix(text, "// @") {
				matches := annotationReg.FindStringSubmatch(text)
				if len(matches) < 2 {
					continue
				}

				annotation := matches[1]

				switch annotation {
				case AnnotationSource, AnnotationInject:
					var injectType string
					var key string

					if len(matches) >= 3 && len(matches[2]) > 0 {
						injectType = InjectByKey
						key = matches[2]
					} else {
						injectType = InjectByType
					}

					item = &AnnotationItem{
						FilePath:       p,
						InjectType:     injectType,
						AnnotationType: annotation,
						Key:            key,
						PackageName:    packageName,
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
			if item.InjectType == InjectByType {
				item.Key = typeName
			}

			switch item.AnnotationType {
			case AnnotationSource:
				key := item.Key
				if _, ok := sourceMap[key]; ok {
					log.Panicf(
						"duplicate source, key: %s, inject type: %s",
						key,
						item.InjectType,
					)
				}

				sourceMap[key] = item
			case AnnotationInject:
				targets = append(targets, item)
			}

			reset()
			continue
		}

		return nil
	})

	cfg = &Config{
		RootPath:  rootPath,
		SourceMap: sourceMap,
		Targets:   targets,
	}

	packageCfg := CreatePackageConfig(cfg)
	cfg.PackageConfig = packageCfg

	return
}
