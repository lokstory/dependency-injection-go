package generator

const ContractTemplate = `package contract

import (
	"reflect"
	"unsafe"
)

type IManager interface {
	SourcePointer(key string) unsafe.Pointer
	SourceValue(key string) reflect.Value
	InjectByValue(sourceKey string, targetValue reflect.Value)
}
`
