package helpers

import (
	"context"
	domainApp "github.com/GleisonEm/bot-claudio-zap-golang/domains/app"
	"mime/multipart"
	"time"
)

func SetAutoConnectAfterBooting(service domainApp.IAppService) {
	time.Sleep(2 * time.Second)
	_ = service.Reconnect(context.Background())
}

func MultipartFormFileHeaderToBytes(fileHeader *multipart.FileHeader) []byte {
	file, _ := fileHeader.Open()
	defer file.Close()

	fileBytes := make([]byte, fileHeader.Size)
	_, _ = file.Read(fileBytes)

	return fileBytes
}

func AudioBytesToFileBytes(audioBytes []byte, filename string) ([]byte, error) {
    fileHeader := &multipart.FileHeader{
        Filename: filename,
        Size:     int64(len(audioBytes)),
    }

    file, err := fileHeader.Open()
    if err != nil {
        return nil, err
    }
    defer file.Close()

    fileBytes := make([]byte, fileHeader.Size)
    _, err = file.Read(fileBytes)
    if err != nil {
        return nil, err
    }

    return fileBytes, nil
}
