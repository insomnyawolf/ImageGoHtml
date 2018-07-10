package main

import "runtime"

const ()

var (
	threads             = runtime.NumCPU()
	thumbnailSize  uint = 128
	alphaPendiente      = float64(1) / float64(65535)
)
