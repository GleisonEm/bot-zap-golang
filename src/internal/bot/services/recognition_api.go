package services

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	DomainBot "github.com/GleisonEm/bot-claudio-zap-golang/domains/bot/structs"
)

func SendToRecogntionApiAudioFile(filePath string) (*DomainBot.ResponseRecognitionApi, error) {
	var requestBody bytes.Buffer

	// Cria um escritor multipart
	multiPartWriter := multipart.NewWriter(&requestBody)

	// Abre o arquivo
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Cria um novo form-file com o nome "audio" para o arquivo
	fileWriter, err := multiPartWriter.CreateFormFile("audio", filePath)
	if err != nil {
		return nil, err
	}

	// Copia o arquivo para o form-file
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	// Fecha o escritor multipart (isso escreve o boundary final)
	err = multiPartWriter.Close()
	if err != nil {
		return nil, err
	}

	// Cria a solicitação POST
	req, err := http.NewRequest("POST", "http://192.168.0.106:5000/extractTextByAudio", &requestBody)
	if err != nil {
		return nil, err
	}

	// Define o Content-Type para multipart/form-data e inclui o boundary
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	// Envia a solicitação
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Decodifica a resposta
	var response DomainBot.ResponseRecognitionApi
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
