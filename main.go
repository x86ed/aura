/*
Copyright © 2022 Adam Siegel incoming@dronestrike.us
*/
//go:generate ./gen.sh
package main

import (
	_ "embed"
	"fmt"

	"github.com/x86ed/aura/cmd"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

//go:embed aura.txt
var aura string

//go:embed ver.txt
var ver string

func main() {
	fmt.Println(aura)
	fmt.Println(ver)
	cmd.Execute()
}
