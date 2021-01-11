package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {

	// pwd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// samplePath := filepath.Join(pwd, "sample")
	// sampleFmt := "%s/%s"

	// fmt.Println(pwd)
	// fmt.Println(samplePath)
	// fmt.Printf(sampleFmt, samplePath, "clockworks.jpg")

	// open "test.jpg"
	file, err := os.Open("sample_images/clockworks.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 200 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	err = jpeg.Encode(out, m, nil)
	if err != nil {
		log.Fatal(err)
	}
}
