package main

import (
	"fmt"
	"image"
)

type enc struct {
	img   image.Image
	y     int
	width int
	ch    chan []byte
}

type lane struct {
	r    []byte
	g    []byte
	b    []byte
	a    []byte
	line []byte
}

func encode(e *enc) []byte {
	buffer := []byte("")
	colspan := 1

	last := color2RGBAHex(e.img.At(0, e.y).RGBA()).RGB()

	for x := 1; x < e.width; x++ {
		col := color2RGBAHex(e.img.At(x, e.y).RGBA()).RGB()
		if col != last || x == (e.width-1) {
			buffer = append(buffer, fmt.Sprintf("d(\"%v\",%v,%v,%v);", col, e.y, x, (x-colspan))...)
			colspan = 1
			last = col
		} else {
			colspan++
		}
	} //Xloop
	//fmt.Println(string(buffer))
	//ReadLine()
	return buffer
}
