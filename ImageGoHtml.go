package main

import (
	"bufio"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/nfnt/resize"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const ()

var (
	threads             = runtime.NumCPU()
	thumbnailSize  uint = 256
	alphaPendiente      = float64(1) / float64(65535)
)

func init() {
	/*
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
		image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	*/
}

func main() {

	fmt.Println("Testing Build")

	args := os.Args[1:]
	argsLenght := len(args)

	if argsLenght == 0 {

		fmt.Println("no value input. Default test value")

		dataIn := "test.jpg"
		idecode(dataIn)

	} else {

		for _, dataIn := range args {
			idecode(dataIn)
		}
	}
}

func idecode(dataIn string) {

	file := openFile(dataIn)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Decode: " + err.Error())
		return
	}

	img = resize.Thumbnail(thumbnailSize, thumbnailSize, img, resize.Lanczos3)

	imgc := img.Bounds()

	height := imgc.Max.Y
	width := imgc.Max.X

	portion := height / threads
	threadLines := portion * threads
	extra := height - threadLines
	fmt.Printf("Image: %v Height: %v Portion: %v Extra: %v\n", file.Name(), height, portion, extra)

	targetY := 0
	startY := 0

	ch := make(chan []byte, threads)

	for thread := 0; thread < threads; thread++ {
		targetY += threadLines

		if extra > 0 {
			targetY++
			extra--
		}

		go encode(img, startY, targetY, width, ch)

		startY = targetY
	}

	data := []byte("<html><body><table border=\"0\" cellpadding=\"1\" cellspacing=\"0\">" + "\n")
	for thread := 0; thread < threads; thread++ {
		data = append(data, <-ch...)
	}
	data = append(data, "</table></body></html>"...)

	result := file.Name() + ".html"
	if exists(result) {
		os.Remove(result)
	}
	ioutil.WriteFile(result, data, 7777)
}

func encode(img image.Image, startY, targetY, targetX int, data chan []byte) {
	buffer := []byte("")
	for y := startY; y < targetY; y++ {
		colspan := 1
		last := ""
		buffer = append(buffer, "<tr>\n"...)
		for x := 0; x < targetX; x++ {
			col := color2RGBAHex(img.At(x, y).RGBA()).RGB()
			if col != last || x == targetX {
				if colspan > 1 {
					buffer = append(buffer, "<td colspan=\""+strconv.Itoa(colspan)+"\" bgcolor=\""+last+"\"></td>"...)
				} else {
					if last != "" {
						buffer = append(buffer, "<td bgcolor=\""+last+"\"></td>"...)
					} else {
						//Do nothing
					}
				}
				colspan = 1
			} else {
				colspan++
			}
			last = col
		} //Xloop
		buffer = append(buffer, "\n</tr>"...)
	} //Yloop
	data <- buffer
}

//Check if file or directiry exist
//If exists returns true
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func openFile(dataIn string) *os.File {
	thing, err := os.Open(dataIn)
	if err != nil {
		log.Fatal(err)
	}
	return thing
}

type colorHEX struct {
	R string
	G string
	B string
	A string
}

func (c colorHEX) RGBA() string {
	return fmt.Sprintf("#%v,%v,%v,%v", c.R, c.G, c.B, c.A)
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

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
