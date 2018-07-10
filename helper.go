package main

import (
	"bufio"
	"log"
	"os"
)

//Exists Check if file or directiry exist
//If exists returns true
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//ReadLine gets coinsole input into string
func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

//OpenFile loads a file into memory
func OpenFile(dataIn string) *os.File {
	thing, err := os.Open(dataIn)
	if err != nil {
		log.Fatal(err)
	}
	return thing
}
