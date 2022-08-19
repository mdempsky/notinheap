// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package notinheap exists.
package notinheap

// Type may be embedded into structs to designate that they should not
// be allocated on the Go heap. For example:
//
//	type MyType struct {
//	    _ notinheap.Type
//	    ...
//	}
type Type struct{ _ nih }

// nih is a separate type so that it can reliably be detected even if
// users write `type Foo notinheap.Type`.
type nih struct{}
