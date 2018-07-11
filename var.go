package main

import (
	"runtime"
)

var (
	threads            = runtime.NumCPU()
	thumbnailSize uint = 512
	coloredOutput      = false
)
