package main

import (
	"image"
	"strconv"
)

type enc struct {
	img   image.Image
	y     int
	width int
	ch    chan []byte
	done  chan struct{}
}

func encode(e *enc) {
	buffer := []byte("<tr>")
	colspan := 1

	last := color2RGBAHex(e.img.At(0, e.y).RGBA()).RGB()

	for x := 1; x < e.width; x++ {

		col := color2RGBAHex(e.img.At(x, e.y).RGBA()).RGB()

		if col != last || x == (e.width-1) {
			if colspan == 1 {
				buffer = append(buffer, "<td bgcolor=\""+last+"\"/>"...)
			} else {
				buffer = append(buffer, "<td colspan=\""+strconv.Itoa(colspan)+"\" bgcolor=\""+last+"\"/>"...)
			}
			colspan = 1
			last = col
		} else {
			colspan++
		}
	} //Xloop
	buffer = append(buffer, "</tr>"...)
	e.done <- struct{}{}
	e.ch <- buffer
}
