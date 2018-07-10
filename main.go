package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/nfnt/resize"
)

func main() {

	fmt.Println("Testing Build")

	args := os.Args[1:]
	argsLenght := len(args)

	if argsLenght == 0 {

		go fmt.Println("no value input. Default test value")

		dataIn := "test.jpg"
		idecode(dataIn)

	} else {

		for _, dataIn := range args {
			idecode(dataIn)
		}

	}
}

func idecode(dataIn string) {
	file := OpenFile(dataIn)
	defer file.Close()

	fmt.Println("Decoding file")

	img := decode(file)

	fmt.Println("Resizing file")
	img = resize.Thumbnail(thumbnailSize, thumbnailSize, img, resize.Lanczos3)

	imgc := img.Bounds()

	height := imgc.Max.Y
	width := imgc.Max.X

	fmt.Printf("Image: %v Width: %v Height: %v \n", file.Name(), width, height)

	var ch []chan []byte
	done := make(chan struct{}, height)

	fmt.Println("Working...")

	for y := 0; y < height; y++ {

		ch = append(ch, make(chan []byte))

		e := enc{
			img:   img,
			y:     y,
			width: width,
			ch:    ch[y],
			done:  done,
		}
		if y > threads {
			<-done
		}

		go encode(&e)

	} //Yloop

	data := []byte("<html><body><table border=\"0\" cellpadding=\"1\" cellspacing=\"0\"><tbody>" + "\n")

	for _, c := range ch {
		data = append(data, <-c...)
	}

	data = append(data, "</tbody></table></body></html>"...)

	result := file.Name() + ".html"
	if Exists(result) {
		os.Remove(result)
	}
	ioutil.WriteFile(result, data, 7777)
}

func decode(file *os.File) image.Image {
	ext := path.Ext(file.Name())
	ext = strings.TrimPrefix(ext, ".")
	ext = strings.ToLower(ext)
	var img image.Image
	var err error
	switch ext {
	case "png":
		img, err = png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	case "jpg":
	case "jpeg":
		img, err = jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	case "gif":
		img, err = gif.Decode(file)
		if err != nil {

		}
	default:
		fmt.Println("default" + ext)
		img, _, err = image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

	}
	return img
}
