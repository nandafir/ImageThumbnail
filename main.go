package main

import (
	"fmt"
	"image"
	"image/color"
	"runtime"

	"github.com/disintegration/imaging"
)

func main() {

	for i := 1; i < 8; i++ {

		nda := i * 100
		err := GenerateThumbnail("sneakers.jpg", nda)
		fmt.Println(err)
	}

}

func GenerateThumbnail(files string, size int) error {
	// use all CPU cores for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// input files
	img, err := imaging.Open(files)
	if err != nil {
		return err
	}

	//get proportional size
	width, height, pointX, pointY := GetImageSize(img, size)

	imgContent := imaging.Thumbnail(img, width, height, imaging.CatmullRom)

	// create a new blank image
	squarePixel := size
	dst := imaging.New(squarePixel, squarePixel, color.White)

	// paste thumbnails into the new image side by side
	// x : horizontal
	// y : vertical
	dst = imaging.Paste(dst, imgContent, image.Pt(pointX, pointY))

	// save the combined image to file
	fileName := fmt.Sprintf("images/01_%v_%v", size, files)
	err = imaging.Save(dst, fileName)
	if err != nil {
		return err
	}

	return nil
}

func GetImageSize(img image.Image, origin int) (int, int, int, int) {

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
			cal := float64(width) / float64(height) * float64(origin)
			width = int(cal)

			height = origin
			pointX = (height - width) / 2

		} else { //landscape

			//get size ratio
			cal := float64(height) / float64(width) * float64(origin)
			height = int(cal)
			width = origin

			pointY = (width - height) / 2
		}

	}

	return width, height, pointX, pointY
}
