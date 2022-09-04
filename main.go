/*
Copyright © 2022 Adam Siegel incoming@dronestrike.us
*/
package main

import (
	_ "embed"
	"fmt"

	"github.com/x86ed/aura/cmd"
)

//go:embed aura.txt
var aura string

func main() {
	fmt.Println(aura)
	cmd.Execute()
}
