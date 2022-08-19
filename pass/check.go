// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// The check command runs the notinheap analyzer.
package main

import (
	"github.com/mdempsky/notinheap/pass"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(pass.Analyzer) }
