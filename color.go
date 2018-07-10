package main

import (
	"fmt"
	"strconv"
)

type colorHEX struct {
	R string
	G string
	B string
	A string
}

func (c colorHEX) RGBA() string {
	return fmt.Sprintf("#%v%v%v%v", c.R, c.G, c.B, c.A)
}

func (c colorHEX) RGB() string {
	return fmt.Sprintf("#%v%v%v", c.R, c.G, c.B)
}

func color2RGBAHex(r, g, b, a uint32) *colorHEX {
	cl := &colorHEX{}
	cl.R = uint2col(r)
	cl.G = uint2col(g)
	cl.B = uint2col(b)
	cl.A = alpha(a)
	return cl
}

func alpha(color uint32) string {
	perc := alphaPendiente * float64(color)
	return strconv.FormatFloat(perc, 'f', 1, 64)
}

func uint2col(color uint32) string {
	col := strconv.FormatUint(uint64(color), 16)
	if last := len(col) - 2; last >= 0 {
		col = col[:last+1]
		col = col[:last]
	}
	if len(col) < 1 {
		col = "00"
	} else if len(col) < 2 {
		col = "0" + col
	}
	return col
}
