package main

import (
	"bufio"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/nfnt/resize"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var (
	wg             sync.WaitGroup
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

	go fmt.Println("Testing Build")

	args := os.Args[1:]
	argsLenght := len(args)

	if argsLenght == 0 {

		go fmt.Println("no value input. Default test value")

		dataIn := "test.jpg"
		idecode(dataIn)

	} else {

		for _, dataIn := range args {

			go fmt.Println(dataIn)
			idecode(dataIn)

		}
	}
}

func idecode(dataIn string) {

	file := openFile(dataIn)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		go fmt.Println("Decode: " + err.Error())
		return
	}

	img = resize.Thumbnail(thumbnailSize, thumbnailSize, img, resize.Lanczos3)

	imgc := img.Bounds()

	buffer := "<html><body><table border=\"0\" cellpadding=\"1\" cellspacing=\"0\"><tr>"

	for y := 0; y < imgc.Max.Y; y++ {
		for x := 0; x < imgc.Max.X; x++ {
			hex := color2RGBAHex(img.At(x, y).RGBA())
			buffer += "<td bgcolor=\"#" + hex.RGB() + "\"></td>"
		}
		buffer += "</tr><tr>"
	}
	buffer += "</tr></table></body></html>"

	bs := []byte(buffer)

	result := file.Name() + ".html"
	if exists(result) {
		os.Remove(result)
	}

	ioutil.WriteFile(result, bs, 7777)

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
	go fmt.Println(thing.Name())

	return thing
}

/*
	wg.Add(1)
	defer wg.Done()
	wg.Wait()
*/

type colorHEX struct {
	R string
	G string
	B string
	A string
}

func (c colorHEX) RGBA() string {
	return fmt.Sprintf("%v%v%v%v", c.R, c.G, c.B, c.A)
}

func (c colorHEX) RGB() string {
	return fmt.Sprintf("%v%v%v", c.R, c.G, c.B)
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
	return col
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
