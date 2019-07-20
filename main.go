package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	"github.com/teris-io/shortid"
)

/*Resize image to 182x268
First we get a random size image. We need to crop it with the correct ratio we want.
Then we resize it. The first step is needed because if not we will get a stretch image.*/

func resizeImg(width, height uint) {

	// open image
	file, err := os.Open("public/3.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	//The default crop use the specified dimension, but it is possible to use Width and Heigth as a ratio instead.
	//  In this case, the resulting image will be as big as possible to fit the asked ratio from the anchor position.
	m, err := cutter.Crop(img, cutter.Config{
		Width:   int(width),
		Height:  int(height),
		Mode:    cutter.Centered,
		Options: cutter.Ratio,
	})
	// now we already crop the image.
	// time to resize.
	m1 := resize.Resize(width, height, m, resize.Lanczos3)
	// create unique name for new image.
	id, err := shortid.Generate()
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("public/ " + id + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m1, nil)

}
func main() {
	resizeImg(182, 268)
}
