/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/x86ed/aura/progbar"
)

var basic bool

// barCmd represents the bar command
var barCmd = &cobra.Command{
	Use:   "bar",
	Short: "Displays a loading bar",
	Long:  `previews various examples of loading bars`,
	Run: func(cmd *cobra.Command, args []string) {
		basic := progbar.BasicProg{}
		basic.New()
		basic.Capacity = 500
		for basic.Count = 0; basic.Count < basic.Capacity+1; basic.Count++ {
			time.Sleep(time.Millisecond * 17)
			basic.Draw()
		}
	},
}

func init() {
	rootCmd.AddCommand(barCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// barCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	barCmd.Flags().BoolVarP(&basic, "basic", "b", false, "Basic Bar chart")
}
