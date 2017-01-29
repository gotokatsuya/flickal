package flickal

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	flickerAPIEndpoint = "https://api.flickr.com/services/rest/?"
)

func buildURL(apiKey string, params map[string]string) (string, error) {
	urlPath, err := url.Parse(flickerAPIEndpoint)
	if err != nil {
		return "", err
	}
	vals := url.Values{}
	for k, v := range params {
		vals.Add(k, v)
	}
	urlPath.RawQuery = vals.Encode()
	return urlPath.String(), nil
}

const (
	flickerSearchPhotoMethod = "flickr.photos.search"
)

func RequestSearchPhoto(apiKey, search string, perPage int, page int) ([]byte, error) {
	parameters := map[string]string{
		"method":   flickerSearchPhotoMethod,
		"api_key":  apiKey,
		"text":     search,
		"per_page": strconv.Itoa(perPage),
		"page":     strconv.Itoa(page),
		"format":   "json",
	}
	builtURL, err := buildURL(apiKey, parameters)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Get(builtURL)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
