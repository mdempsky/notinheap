// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pass provides an analyzer for misuse of notinheap.Type.
package pass

import (
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

var Analyzer = &analysis.Analyzer{
	Name:     "notinheap",
	Doc:      "check for heap allocations of notinheap.Type",
	Requires: []*analysis.Analyzer{buildssa.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	fact := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	for _, fn := range fact.SrcFuncs {
		runFunc(pass, fn)
	}
	return nil, nil
}

func runFunc(pass *analysis.Pass, fn *ssa.Function) {
	for _, clo := range fn.AnonFuncs {
		runFunc(pass, clo)
	}

	for _, block := range fn.Blocks {
		runBlock(pass, block)
	}
	if block := fn.Recover; block != nil {
		runBlock(pass, block)
	}
}

func runBlock(pass *analysis.Pass, block *ssa.BasicBlock) {
	for _, instr := range block.Instrs {
		switch instr := instr.(type) {
		case *ssa.Alloc:
			if instr.Heap {
				checkType(pass, instr, instr.Type().Underlying().(*types.Pointer).Elem())
			}
		case *ssa.MakeSlice:
			checkType(pass, instr, instr.Type().Underlying().(*types.Slice).Elem())
		}
	}
}

func checkType(pass *analysis.Pass, instr ssa.Instruction, typ types.Type) {
	if nih(typ) {
		pass.Reportf(instr.Pos(), "heap allocation of %v: %v", types.TypeString(typ, types.RelativeTo(pass.Pkg)), instr)
	}
}

// nih reports whether typ is or contains notinheap.Type.
func nih(typ types.Type) bool {
	switch typ := typ.(type) {
	case *types.Named:
		if obj := typ.Obj(); obj.Name() == "nih" && obj.Pkg().Path() == "github.com/mdempsky/notinheap" {
			return true
		}
		return nih(typ.Underlying())

	case *types.Array:
		return nih(typ.Elem())
	case *types.Struct:
		for i := 0; i < typ.NumFields(); i++ {
			if nih(typ.Field(i).Type()) {
				return true
			}
		}
	}

	return false
}
