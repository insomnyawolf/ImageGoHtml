package main

import (
	"fmt"

	ansi "github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar"
)

type progressStat struct {
	current int
	total   int
	height  int
}

//NewProgressBar creates a new progressbar
func NewProgressBar(progress *progressStat) *progressbar.ProgressBar {
	if coloredOutput {
		return progressbar.NewOptions(
			progress.height,
			progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
			progressbar.OptionSetWidth(100),
			//progressbar.OptionSetBytes(10000),
			progressbar.OptionSetRenderBlankState(true),
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionSetDescription(
				fmt.Sprintf("[cyan][%v/%v][reset] Encoding...", progress.current, progress.total),
			),
			progressbar.OptionSetTheme(
				progressbar.Theme{
					Saucer: "[cyan]#[reset]",
					//SaucerHead:    ">",
					SaucerPadding: "-",
					BarStart:      ">",
					BarEnd:        "<",
				},
			),
		)
	}
	return progressbar.NewOptions(
		progress.height,
		progressbar.OptionSetWidth(100),
		//progressbar.OptionSetBytes(10000),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetDescription(
			fmt.Sprintf("[%v/%v] Encoding...", progress.current, progress.total),
		),
		progressbar.OptionSetTheme(
			progressbar.Theme{
				Saucer: "#",
				//SaucerHead:    ">",
				SaucerPadding: "-",
				BarStart:      ">",
				BarEnd:        "<",
			},
		),
	)
}
