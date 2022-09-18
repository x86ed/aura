/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/x86ed/aura/modal"
)

// modalCmd represents the modal command
var modalCmd = &cobra.Command{
	Use:   "modal",
	Short: "renders a modal",
	Long:  `renders a modal with user specified options`,
	Run: func(cmd *cobra.Command, args []string) {
		g := modal.ModalGlyph{
			LUC: "╔",
			RUC: "╗",
			RDC: "╝",
			LDC: "╚",
			L:   "║",
			R:   "║",
			U:   "═",
			D:   "═",
			F:   " ",
		}
		m := modal.Modal{
			GFX: g,
		}
		m.Dims.X = 10
		m.Dims.Y = 5
		m.Offset.X = 2
		m.Offset.Y = 2
		m.IsOffset = true
		m.Draw()
	},
}

func init() {
	rootCmd.AddCommand(modalCmd)
}
