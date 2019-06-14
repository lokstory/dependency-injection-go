package contract

import "unsafe"

type IManager interface {
	GetPointer(string) unsafe.Pointer
}
