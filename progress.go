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
	if cnf.IsColored {
		return progressbar.NewOptions(
			progress.height,
			progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
			//progressbar.OptionSetWidth(100),
			//progressbar.OptionSetBytes(10000),
			progressbar.OptionSetRenderBlankState(true),
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionSetDescription(
				fmt.Sprintf("[%v][%v/%v][reset] Encoding...", cnf.Color, progress.current, progress.total),
			),
			progressbar.OptionSetTheme(
				progressbar.Theme{
					Saucer: fmt.Sprintf("[%v]#[reset]", cnf.Color),
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
		//progressbar.OptionSetWidth(100),
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
