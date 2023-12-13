package chroma

import (
	"avec_moi_with_us_api/api/utils"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Repository struct {
	baseUrl string
	num     string
}

func NewRepository() *Repository {
	return &Repository{baseUrl: utils.GetChromaClientUrl(), num: utils.GetChromaNum()}
}
func (r Repository) Search(content string) ([]string, error) {
	var result []string
	encodedContent := url.QueryEscape(content)
	response, err := http.Get(r.baseUrl + "/api/v1/movie/search?content=" + encodedContent + "&num=" + r.num)

	if err != nil {
		return result, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			utils.LogError(utils.LogData{Event: "Failed to close chroma db response", User: "system", Details: err.Error()})
		}
	}(response.Body)
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil

}
