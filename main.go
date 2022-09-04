/*
Copyright © 2022 Adam Siegel incoming@dronestrike.us
*/
package main

import (
	"fmt"

	"github.com/x86ed/aura/cmd"
)

func main() {
	fmt.Println("\u001b[6n")
	cmd.Execute()
}
