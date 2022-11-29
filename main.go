package main

import (
	"image"
	"image/color"
	"runtime"

	"github.com/disintegration/imaging"
)

func main() {
	// use all CPU cores for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// input files
	files := "sneakers.jpg"
	img, err := imaging.Open(files)
	if err != nil {
		panic(err)
	}

	//get proportional size
	width, height, pointX, pointY := getImageSize(img, 500)

	imgContent := imaging.Thumbnail(img, width, height, imaging.CatmullRom)

	// create a new blank image
	squarePixel := 500
	dst := imaging.New(squarePixel, squarePixel, color.White)

	// paste thumbnails into the new image side by side
	// x : horizontal
	// y : vertical
	dst = imaging.Paste(dst, imgContent, image.Pt(pointX, pointY))

	// save the combined image to file
	err = imaging.Save(dst, "01_thumbnail-"+files)
	if err != nil {
		panic(err)
	}
}

func getImageSize(img image.Image, origin int) (int, int, int, int) {

	var height, width, pointX, pointY int

	i := img.Bounds()
	width = i.Dx()
	height = i.Dy()

	if height == width {
		width = origin
		height = origin

	} else {

		//get default max:pixel
		if height > width { //potrait

			//get size ratio
			if width > origin {
				cal := float64(width) / float64(height) * float64(origin)
				width = int(cal)
			}

			height = origin
			pointX = (height - width) / 2

		} else { //landscape

			//get size ratio
			if height > origin {
				cal := float64(height) / float64(width) * float64(origin)
				height = int(cal)
			}

			width = origin
			pointY = (width - height) / 2
		}

	}

	return width, height, pointX, pointY
}
