/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/x86ed/aura/block"
	"github.com/x86ed/aura/color"
	"github.com/x86ed/aura/modal"
)

// modalCmd represents the modal command
var modalCmd = &cobra.Command{
	Use:   "modal",
	Short: "renders a modal",
	Long:  `renders a modal with user specified options`,
	Run: func(cmd *cobra.Command, args []string) {
		c := block.Block{
			Content: color.BgBlue("this is a\ntest.\n\nhello.hello.\nthis is another test\neyooooooooooooo"),
		}
		c2 := block.Block{
			Content: color.BgRed("this is a\ntest."),
		}
		c.Next = &c2
		g := modal.FrameDouble
		m := modal.Modal{
			GFX:     g,
			Content: &c,
			Padding: 1,
		}
		m.Dims.X = 20
		m.Dims.Y = 20
		m.Offset.X = 2
		m.Offset.Y = 5
		m.IsOffset = true
		m.Draw()
	},
}

func init() {
	rootCmd.AddCommand(modalCmd)
}
