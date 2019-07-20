package img

import (
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	"github.com/teris-io/shortid"
)

// Insert contact data
func Insert(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	fmt.Println(filename)
	url := resizeImg(182, 268, filename)
	c.JSON(200, gin.H{
		"img": url,
	})
}

/*Resize image to 182x268
First we get a random size image. We need to crop it with the correct ratio we want.
Then we resize it. The first step is needed because if not we will get a stretch image.*/

func resizeImg(width, height uint, imgPath string) string {

	savePath := "public/images/"
	baseURL := "http://localhost:8080/"

	// open image
	file, err := os.Open(imgPath)
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
	// public/images/image_name.jpg
	//Related file path + file name + ext
	relatedOut := savePath + id + ".jpg"
	out, err := os.Create(relatedOut)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m1, nil)
	// remove the original file
	err = os.Remove(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	return baseURL + relatedOut

}
