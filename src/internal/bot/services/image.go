package services

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/chai2010/webp"
	"github.com/google/uuid"
)

func ConvertToWebp(outputPath string, pathMedia string) (string, error) {
	// Open the PNG file
	pngFile, err := os.Open(pathMedia)
	if err != nil {
		return "", err
	}
	defer pngFile.Close()

	// Decode the PNG file to get the image
	img, err := png.Decode(pngFile)
	if err != nil {
		return "", err
	}

	// Check if the image is NRGBA
	nrgba, ok := img.(*image.NRGBA)
	if !ok {
		return "", fmt.Errorf("image is not NRGBA")
	}

	// Encode the image to WebP
	webpData, err := webp.EncodeLosslessRGBA(nrgba)
	if err != nil {
		return "", err
	}

	// Save the WebP image to a file
	outputpathMedia := fmt.Sprintf("%s/%d-%s%s", outputPath, time.Now().Unix(), uuid.NewString(), ".webp")
	webpFile, err := os.Create(outputpathMedia)
	if err != nil {
		return "", err
	}
	defer webpFile.Close()

	_, err = webpFile.Write(webpData)
	if err != nil {
		return "", err
	}

	return outputpathMedia, nil
}
