package p

import (
	"github.com/mdempsky/notinheap"
)

type x struct{ notinheap.Type }

func f() *x { return new(x) }

func h() *x { return &x{} }

var t [4]x

func g(n int) any { return make([]x, n) }
