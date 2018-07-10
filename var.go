package main

import (
	"runtime"
)

var (
	threads             = runtime.NumCPU()
	thumbnailSize  uint = 2500
	alphaPendiente      = float64(1) / float64(65535)
)
