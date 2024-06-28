package services

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"time"

	"github.com/chai2010/webp"
	"github.com/google/uuid"
)

func ConvertPngToWebp(outputPath string, pathMedia string) (string, error) {
	// Open the PNG file
	pngFile, err := os.Open(pathMedia)
	if err != nil {
		fmt.Println("passei aq pathMedia")
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

func ConvertJpegToWebp(outputPath string, pathMedia string) (string, error) {
	// Open the JPEG file
	jpegFile, err := os.Open(pathMedia)
	if err != nil {
		fmt.Println("Failed to open JPEG file:", pathMedia)
		return "", err
	}
	defer jpegFile.Close()

	// Decode the JPEG file to get the image
	img, err := jpeg.Decode(jpegFile)
	if err != nil {
		return "", err
	}

	// Convert the image to NRGBA
	bounds := img.Bounds()
	nrgba := image.NewNRGBA(bounds)
	draw.Draw(nrgba, bounds, img, bounds.Min, draw.Src)

	// Encode the image to WebP
	webpData, err := webp.EncodeLosslessRGBA(nrgba)
	if err != nil {
		return "", err
	}

	// Save the WebP image to a file
	outputPathMedia := fmt.Sprintf("%s/%d-%s%s", outputPath, time.Now().Unix(), uuid.NewString(), ".webp")
	webpFile, err := os.Create(outputPathMedia)
	if err != nil {
		return "", err
	}
	defer webpFile.Close()

	_, err = webpFile.Write(webpData)
	if err != nil {
		return "", err
	}

	return outputPathMedia, nil
}
