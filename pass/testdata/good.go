package p

import (
	"unsafe"

	"github.com/mdempsky/notinheap"
)

type x struct{ notinheap.Type }

func f(p uintptr) *x { return (*x)(unsafe.Pointer(p)) }

func g(p *x) { *p = x{} }
