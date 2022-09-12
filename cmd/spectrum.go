/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/x86ed/aura/color"
)

var p256 bool
var rgb bool

// spectrumCmd represents the spectrum command
var spectrumCmd = &cobra.Command{
	Use:   "spectrum",
	Short: "Displays various spectrums",
	Long: `Displays various specrums supported by the X11 terminal:
	16 Color - System color pallete. This is set by the terminal preferences but matches the default Low and Hi 8 color palletes

	256 Color - 216 color color cube pallete as defined by the 256 color pallete

	24 Shade Grey - greyscale from the 256 color pallete
	
	RGB Color - 24 bit RGB color
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(genSpectrums())
	},
}

func genSpectrums() string {
	var ss string
	if p256 {
		ss += "             Standard Color\n                (system)\n"
		for c := 0; c < 16; c++ {
			ss += color.Bg256(c, fmt.Sprintf(" %3d ", c))
			if c == 7 {
				ss += "\n                Hi Color\n                (system)\n"
			}
		}
		ss += "\n                216 Color\n     "
		for i := 16; i < 232; i++ {
			ss += color.Bg256(i, fmt.Sprintf(" %3d ", i))
			if i%3 == 0 && i%2 == 1 {
				ss += "\n     "
			}
		}
		ss += "\n                Greyscale\n    "
		for g := 232; g < 256; g++ {
			l := " "
			if g == 232 || g == 255 {
				l = fmt.Sprintf(" %3d ", g)
			}
			ss += color.Bg256(g, l)
		}
		return ss
	}
	return ""
}

func init() {
	rootCmd.AddCommand(spectrumCmd)
	spectrumCmd.Flags().BoolVarP(&p256, "256", "x", false, "Display 256 color spectrum")
	spectrumCmd.Flags().BoolVarP(&rgb, "rgb", "c", false, "Display RGB 24 bit spectrum")
}
