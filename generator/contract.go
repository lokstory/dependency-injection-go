package generator

const ContractTemplate = `
package contract

import "unsafe"

type IManager interface {
	GetPointer(string) unsafe.Pointer
}
`
