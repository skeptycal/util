package main

// package testing setup
// Copyright (c) 2020 Michael Treanor
// MIT License

import (
	"fmt"
	"strings"

	"github.com/skeptycal/anansi"
)

func stringBuilderCheck() {
	var sb strings.Builder
	sb.Reset()
	defer sb.Reset()

	for i := 0; i < 10; i++ {
		sb.WriteString("a")
	}

	fmt.Println(sb.String())
}

func main() {
	anansi.Sample()

}
