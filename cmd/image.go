/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	"github.com/x86ed/aura/color"
	"github.com/x86ed/aura/img"
	"github.com/x86ed/aura/screen"
	"github.com/x86ed/aura/util"
)

var qr bool
var jpg bool
var gif bool
var png bool
var value string

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "displays images in the terminal",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if qr {
			QR := img.QR{
				Size:     21,
				Recovery: qrcode.Low,
				FGColor:  color.BgBlack(),
				BGColor:  color.BgWhite(),
			}
			txt := "https://github.com/x86ed/aura"
			if len(value) > 0 {
				txt = value
			}
			QR.QRFromText(txt)
		}
		if png {
			txt := "./img/kali.png"
			if len(value) > 0 {
				txt = value
			}
			var scr screen.S

			scr.GetDims()
			i := img.Img{
				FilePath: txt,
				File:     img.PNG,
				Fit:      img.FitWidth,
			}
			i.Dims = util.Coord{X: scr.Width, Y: scr.Height}
			i.Draw()
		}
		if jpg {
			txt := "./img/image.jpeg"
			if len(value) > 0 {
				txt = value
			}
			var scr screen.S

			scr.GetDims()
			i := img.Img{
				FilePath: txt,
				File:     img.JPG,
				Fit:      img.FitHeight,
			}
			i.Dims = util.Coord{X: scr.Width, Y: scr.Height}
			i.Draw()
		}
		if gif {
			txt := "./img/v.gif"
			if len(value) > 0 {
				txt = value
			}
			var scr screen.S

			scr.GetDims()
			i := img.Gif{
				FilePath: txt,
				Fit:      img.FitHeight,
			}
			i.Dims = util.Coord{X: scr.Width, Y: scr.Height}
			i.Draw()
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.Flags().BoolVarP(&qr, "QR", "q", false, "draws a QR Code that points to https://github.com/x86ed/aura or a value specified by the --value (-v) flag")
	imageCmd.Flags().BoolVarP(&jpg, "jpg", "j", false, "draws a jpg a value specified by the --value (-v) flag")
	imageCmd.Flags().BoolVarP(&gif, "gif", "g", false, "draws a gif a value specified by the --value (-v) flag")
	imageCmd.Flags().BoolVarP(&png, "png", "p", false, "draws a png a value specified by the --value (-v) flag")
	imageCmd.Flags().StringVarP(&value, "value", "v", "", "string value passed in to method (values with spaces need to be contained in quotes)")
}
