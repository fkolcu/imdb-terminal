package util

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const emptyImage = "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"

// RetrieveImage downloads image from the given url and
// returns the base64 encoded image
func RetrieveImage(imageUrl string) string {
	imageFullPath, err := downloadImage(imageUrl)
	if err != nil {
		return emptyImage
	}

	base64Image, err := imageToBase64(imageFullPath)
	if err != nil {
		return emptyImage
	}

	return base64Image
}

func RetrieveEmptyImage() string {
	return emptyImage
}

func downloadImage(imageUrl string) (string, error) {
	tempDir, err := GetTempDir()
	if err != nil {
		return "", err
	}

	destinationPath := filepath.Join(tempDir, "imdb-poster.jpg")

	response, err := http.Get(imageUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	file, err := os.Create(destinationPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return destinationPath, nil
}

func imageToBase64(imagePath string) (string, error) {
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	imageData, err := io.ReadAll(imageFile)
	if err != nil {
		return "", err
	}

	base64String := base64.StdEncoding.EncodeToString(imageData)
	return base64String, nil
}
