package main

import (
	"context"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/nfnt/resize"
	"golang.org/x/sync/semaphore"
)

func init() {
	config()
}

func main() {

	fmt.Println("Testing Build")

	args := os.Args[1:]
	argsLenght := len(args)

	if argsLenght == 0 {
		path, _ := os.Executable()
		fmt.Printf("%v (path to pic)", path)

		ReadLine()

	} else {
		for current, dataIn := range args {
			idecode(dataIn, &progressStat{current: (current + 1), total: argsLenght})
		}
	}
}

func idecode(dataIn string, progress *progressStat) {
	file := OpenFile(dataIn)
	defer file.Close()

	img := decode(file)

	imgc := img.Bounds()
	height := imgc.Max.Y
	width := imgc.Max.X

	if cnf.IsResized {
		if uint(height) > cnf.MaxSize || uint(width) > cnf.MaxSize {
			img = resize.Thumbnail(cnf.MaxSize, cnf.MaxSize, img, resize.Lanczos3)
			imgc = img.Bounds()
			height = imgc.Max.Y
			width = imgc.Max.X
		}
	}

	fmt.Printf("Image: %v Width: %v Height: %v \n", file.Name(), width, height)

	progress.height = height
	bar := NewProgressBar(progress)

	var ch []chan []byte
	ctx := context.TODO()
	sem := semaphore.NewWeighted(int64(threads))

	for y := 0; y < height; y++ {
		sem.Acquire(ctx, 1)
		ch = append(ch, make(chan []byte))
		e := &enc{
			img:   img,
			y:     y,
			width: width,
			ch:    ch[y],
		}
		go func(ec *enc) {
			ans := encode(ec)
			bar.Add(1)
			sem.Release(1)
			ec.ch <- ans
		}(e)
	} //Yloop

	bar.Finish()

	canvasID := "c"

	data := []byte(
		fmt.Sprintf("<html>"+
			"<body>"+
			"<canvas id=\"%v\" width=\"%v\" height=\"%v\"></canvas>"+
			"<script>"+
			"var ct=document.getElementById(\"%v\");"+
			"var c=ct.getContext(\"2d\");"+
			"function d(color , y, startX, endX){"+
			"c.fillStyle='#'+color;"+
			"c.fillRect(startX,y,endX,y);"+
			"}",
			canvasID, width, height, canvasID),
	)
	data = append(data, ""...)
	for _, c := range ch {
		data = append(data, <-c...)
	}

	data = append(data, "</script></body></html>"...)

	fmt.Println("")

	result := file.Name() + ".html"
	if Exists(result) {
		os.Remove(result)
	}
	ioutil.WriteFile(result, data, 7777)
}

//Ways to optimize that?
func decode(file *os.File) image.Image {
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
