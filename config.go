package main

import (
	"fmt"
	"os"

	"github.com/thehowl/conf"
)

type myConf struct {
	//Image Settings
	IsResized bool `description:"if true resize the image"`
	MaxSize   uint `description:"If isResized is true, this will define the longest side size"`
	//Output
	IsColored bool   `description:"if true colors the console output"`
	Color     string `description:"if isColored is true will color the output with this color as example cyan"`
}

func config() {
	config := myConf{}
	err := conf.Load(&config, "options.conf")
	if err == conf.ErrNoFile {
		conf.Export(myConf{
			//Image Settings
			IsResized: true,
			MaxSize:   512,
			//Output
			IsColored: true,
			Color:     "cyan",
		}, "options.conf")
		fmt.Println("options.conf was created sucesfully.")
		os.Exit(1)
	}
	if err != nil {
		panic(err)
	}

	cnf = config
}
