package main

import (
	"runtime"
)

var (
	cnf     myConf
	threads = runtime.NumCPU()
)
