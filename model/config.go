package model

type Config struct {
	RootPath string
	SourceMap map[string]*AnnotationItem
	Targets []*AnnotationItem
}
