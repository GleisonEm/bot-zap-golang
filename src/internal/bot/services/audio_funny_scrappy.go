package services

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"encoding/base64"
	"io/ioutil"

	"github.com/antchfx/htmlquery"
)

type SearchError struct {
	Message string
	Err     error
}

func (e *SearchError) Error() string {
	return e.Message
}

func SearchAudioFunny(name string) (map[string]string, error) {
	if name == "" {
		return nil, &SearchError{Message: "Parâmetro 'name' não especificado na URL"}
	}

	searchURL := fmt.Sprintf("https://www.myinstants.com/pt/search/?name=%s", name)
	doc, err := htmlquery.LoadURL(searchURL)
	if err != nil {
		return nil, &SearchError{Message: "Erro ao analisar HTML", Err: err}
	}

	button := htmlquery.FindOne(doc, "//*[@id='instants_container']/div[1]/div[1]/button")
	if button == nil {
		return nil, &SearchError{Message: "Botão não encontrado"}
	}

	onclick := htmlquery.SelectAttr(button, "onclick")
	if onclick == "" {
		return nil, &SearchError{Message: "Atributo 'onclick' não encontrado"}
	}

	soundPath := extractSoundPath(onclick)
	title := htmlquery.SelectAttr(button, "title")
	if title == "" {
		return nil, &SearchError{Message: "Atributo 'title' não encontrado"}
	}

	response := map[string]string{
		"soundPath": soundPath,
		"title":     title,
	}

	return response, nil
}

func SearchAudioFunnyAll(name string) ([]string, error) {
	if name == "" {
		return nil, &SearchError{Message: "Parâmetro 'name' não especificado na URL"}
	}

	searchURL := fmt.Sprintf("https://www.myinstants.com/pt/search/?name=%s", name)
	doc, err := htmlquery.LoadURL(searchURL)
	if err != nil {
		return nil, &SearchError{Message: "Erro ao analisar HTML", Err: err}
	}

	buttons := htmlquery.Find(doc, "//*[@class='instant-link link-secondary']")
	if buttons == nil {
		return nil, &SearchError{Message: "Botões não encontrados"}
	}

	var titles []string

	for _, button := range buttons {
		title := htmlquery.InnerText(button)
		titles = append(titles, title)
	}

	return titles, nil
}

func SearchAudioFunnyReturnBase64(name string) (map[string]string, error) {
	if name == "" {
		return nil, &SearchError{Message: "Parâmetro 'name' não especificado na URL"}
	}

	searchURL := fmt.Sprintf("https://www.myinstants.com/pt/search/?name=%s", name)
	doc, err := htmlquery.LoadURL(searchURL)
	if err != nil {
		return nil, &SearchError{Message: "Erro ao analisar HTML", Err: err}
	}

	button := htmlquery.FindOne(doc, "//*[@id='instants_container']/div[1]/div[1]/button")
	if button == nil {
		return nil, &SearchError{Message: "Botão não encontrado"}
	}

	onclick := htmlquery.SelectAttr(button, "onclick")
	if onclick == "" {
		return nil, &SearchError{Message: "Atributo 'onclick' não encontrado"}
	}

	soundPath := extractSoundPath(onclick)
	url := fmt.Sprintf("https://www.myinstants.com%s", soundPath)
	base64String, err := getBase64Media(url)
	if err != nil {
		return nil, &SearchError{Message: "Erro ao obter o vídeo em base64", Err: err}
	}

	response := map[string]string{
		"soundBase64": base64String,
	}

	return response, nil
}

func SearchAudioFunnyReturnFile(name string) ([]byte, error) {
	if name == "" {
		return nil, &SearchError{Message: "Parâmetro 'name' não especificado na URL"}
	}

	searchURL := fmt.Sprintf("https://www.myinstants.com/pt/search/?name=%s", name)
	doc, err := htmlquery.LoadURL(searchURL)
	if err != nil {
		return nil, &SearchError{Message: "Erro ao analisar HTML", Err: err}
	}

	button := htmlquery.FindOne(doc, "//*[@id='instants_container']/div[1]/div[1]/button")
	if button == nil {
		return nil, &SearchError{Message: "Botão não encontrado"}
	}

	onclick := htmlquery.SelectAttr(button, "onclick")
	if onclick == "" {
		return nil, &SearchError{Message: "Atributo 'onclick' não encontrado"}
	}

	soundPath := extractSoundPath(onclick)
	url := fmt.Sprintf("https://www.myinstants.com%s", soundPath)

	return downloadAudio(url)
}

func getBase64Media(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Request failed with status code: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	base64String := base64.StdEncoding.EncodeToString(body)

	return base64String, nil
}

func downloadAudio(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func extractSoundPath(onclick string) string {
	format := strings.Replace(onclick, "play(", "", -1)
	splitParts := strings.Split(format, ",")

	return strings.Trim(splitParts[0], "'")
}
