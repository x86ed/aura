/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/x86ed/aura/v2/color"
)

type spectrum struct {
	bg256 [13][13]int
	rgb   [][]int
}

var p256 bool

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
	var spec = spectrum{bg256: [13][13]int{
		{231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231},
		{254, 224, 213, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{252, 217, 212, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{250, 210, 211, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{248, 210, 209, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{246, 203, 209, 0, 0, 83, 0, 0, 0, 0, 0, 0, 0},
		{244, 196, 208, 228, 118, 46, 48, 51, 33, 21, 93, 201, 198},
		{242, 160, 172, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{240, 124, 136, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{238, 88, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{236, 1, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{234, 52, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
	}}

	if p256 {
		for i := 0; i < 13; i++ {
			for j := 0; j < 13; j++ {
				ss += color.Bg256(spec.bg256[i][j], fmt.Sprintf(" %3d ", spec.bg256[i][j]))
			}
			ss += "\n"
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
}
