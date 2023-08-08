package view

import (
	"bytes"
	"encoding/base64"
	"github.com/fkolcu/imdb-terminal/internal/util"
	"github.com/rivo/tview"
	"image/jpeg"
)

func NewImage(base64Image string) *tview.Image {
	if base64Image == "" {
		base64Image = util.RetrieveEmptyImage()
	}

	myImage := tview.NewImage()
	imageBytes, _ := base64.StdEncoding.DecodeString(base64Image)
	photo, _ := jpeg.Decode(bytes.NewReader(imageBytes))
	myImage.SetImage(photo)
	myImage.SetColors(tview.TrueColor)

	return myImage
}
