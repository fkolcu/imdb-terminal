package view

import (
	"bytes"
	"encoding/base64"
	"github.com/rivo/tview"
	"image/jpeg"
)

func NewImage(base64Image string) *tview.Image {
	myImage := tview.NewImage()
	imageBytes, _ := base64.StdEncoding.DecodeString(base64Image)
	photo, _ := jpeg.Decode(bytes.NewReader(imageBytes))
	myImage.SetImage(photo)
	myImage.SetColors(tview.TrueColor)

	return myImage
}
