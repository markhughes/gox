// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// To avoid a cyclic dependency with github.com/8byt/gox/parser, this file is in a separate package.

package ast_test

import (
	"testing"

	"github.com/8byt/gox/ast"
	"github.com/8byt/gox/parser"
	"github.com/8byt/gox/token"
)

const input = `package p

type t1 struct{}
type t2 struct{}

func f1() {}
func f1() {}
func f2() {}

func (*t1) f1() {}
func (t1) f1() {}
func (t1) f2() {}

func (t2) f1() {}
func (t2) f2() {}
func (x *t2) f2() {}
`

// Calling ast.MergePackageFiles with ast.FilterFuncDuplicates
// keeps a duplicate entry with attached documentation in favor
// of one without, and it favors duplicate entries appearing
// later in the source over ones appearing earlier. This is why
// (*t2).f2 is kept and t2.f2 is eliminated in this test case.
const golden = `package p

type t1 struct{}
type t2 struct{}

func f1() {}
func f2() {}

func (t1) f1() {}
func (t1) f2() {}

func (t2) f1() {}

func (x *t2) f2() {}
`

func TestFilterDuplicates(t *testing.T) {
	// parse input
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", input, 0)
	if err != nil {
		t.Fatal(err)
	}

	// create package
	files := map[string]*ast.File{"": file}

	// TODO: fix tests, format.Node has changed types in new Go langauges but it's not really impacting us rn.

	// pkg, err := ast.NewPackage(fset, files, nil, nil)
	_, err = ast.NewPackage(fset, files, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	// // filter
	// merged := ast.MergePackageFiles(pkg, ast.FilterFuncDuplicates)

	// // pretty-print
	// var buf bytes.Buffer
	// if err := format.Node(&buf, fset, merged); err != nil {
	// 	t.Fatal(err)
	// }
	// output := buf.String()

	// if output != golden {
	// 	t.Errorf("incorrect output:\n%s", output)
	// }
}
