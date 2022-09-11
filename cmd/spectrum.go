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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spectrumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spectrumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	spectrumCmd.Flags().BoolVarP(&p256, "256", "x", false, "Display 256 color spectrum")
	spectrumCmd.Flags().BoolVarP(&rgb, "rgb", "c", false, "Display RGB 24 bit spectrum")
}
