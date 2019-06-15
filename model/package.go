package model

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

// PackageConfig
type PackageConfig struct {
	// ItemMap
	ItemMap  map[string]*PackageItem
	// Packages
	Packages []string
	// DirMap
	DirMap   map[string]*PackageItem
}

// PackageItem
type PackageItem struct {
	Path  string
	Alias string
}

func CreatePackageConfig(cfg *Config) *PackageConfig {
	var packages []string
	// key: dir, value: alias
	packageMap := map[string]string{}
	// packageName -> DIR -> id
	aliasMap := map[string]map[string]*PackageItem{}

	itemMap := map[string]*PackageItem{}
	dirMap := map[string]*PackageItem{}

	var setByAnnotationItem = func(item *AnnotationItem) {
		if _, ok := itemMap[item.FilePath]; ok {
			return
		}

		relPath, err := filepath.Rel(cfg.RootPath, item.FilePath)
		if err != nil {
			log.Panic(err)
		}

		dir := filepath.Dir(relPath)
		dir = strings.ReplaceAll(dir, `\`, `/`)
		var packageName string

		if lastID := strings.LastIndex(dir, "/"); lastID >= 0 {
			packageName = dir[lastID+1:]
		} else {
			packageName = dir
		}

		var alias string
		var pathItem *PackageItem

		if _, ok := aliasMap[packageName]; !ok {
			alias = packageName
			pathItem = &PackageItem{
				Path:  dir,
				Alias: alias,
			}
			aliasMap[packageName] = map[string]*PackageItem{}
			aliasMap[packageName][dir] = pathItem
		} else {
			pathItem, ok = aliasMap[packageName][dir]
			if !ok {
				alias = fmt.Sprintf("%s%d", packageName, len(aliasMap[packageName])+1)
				pathItem = &PackageItem{
					Path:  dir,
					Alias: alias,
				}
				aliasMap[packageName][dir] = pathItem
			}
		}

		itemMap[item.FilePath] = pathItem
		dirMap[dir] = pathItem

		if _, ok := packageMap[dir]; !ok {
			packageMap[dir] = alias
			packages = append(packages, dir)
		}
	}

	for _, item := range cfg.SourceMap {
		setByAnnotationItem(item)
	}

	for _, item := range cfg.Targets {
		setByAnnotationItem(item)
	}

	sort.Strings(packages)

	return &PackageConfig{
		ItemMap:  itemMap,
		DirMap:   dirMap,
		Packages: packages,
	}
}
