// Copyright 2021 Michael Treanor. All rights reserved.
// Use of this source code is governed the MIT License
// that can be found in the LICENSE file.

// Package gogen implements automation for Go code generation
// using go generate and templates.
//
// usage example:
//
// $ gogen maketests.go
//
// $ go run maketests.go -output benchmarks.go
//
package gogen

import (
	"fmt"
	"strings"
)


const (
	fmtCopyrightHeader = `// Copyright (c) %s %s. All rights reserved.
// Use of this source code is governed the %s License
// that can be found in the LICENSE file.

`
	fmtCodeGenNotice = "// Code generated by go run %s -output %s; DO NOT EDIT.\n\n"

	fmtPackage       = "package %s"
	fmtMemo          = "// %s"
	fmtSectionCloser = "}\n\n"
	fmtFuncHeader    = "//%s is an automatically generated function.\nfunc %s (%s) %s {\n"
	fmtStructHeader  = "//%s is an automatically generated struct.\ntype %s struct {\n"
	fmtTableHeader   = "//%s is an automatically generated table.\nvar %s %s = %s{\n"
)


func (r *RepoConfig) genPageTemplate() string {
	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteString(r.genCopyrightHeader())
	// todo -
	return sb.String()
}

func (r *RepoConfig) genCopyrightHeader() string {
	return fmt.Sprintf(fmtCopyrightHeader, r.year, r.User.Name, r.license)
}

func (r *RepoConfig) genPackage() string {
	return fmt.Sprintf(fmtPackage, r.name)
}

func (r *RepoConfig) genMemo(memo string) string {
	return fmt.Sprintf(fmtMemo, memo)
}

func genClose() string { return fmtSectionCloser }

func (r *RepoConfig) genFuncHeader(name, args, retvals string) string {
	return fmt.Sprintf(fmtFuncHeader, name, name, args, retvals)
}

func (r *RepoConfig) genStructHeader(name string) string {
	return fmt.Sprintf(fmtStructHeader, name, name)
}

func (r *RepoConfig) genTableHeader(name, vartype, value string) string {
	return fmt.Sprintf(fmtTableHeader, name, name, vartype, value)
}
