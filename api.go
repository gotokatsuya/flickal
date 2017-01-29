package flickal

import "net/http"

func SearchPhotoWithHTTPClient(c *http.Client, apiKey, search string, perPage int, page int) (*Images, error) {
	rawBody, err := RequestSearchPhoto(c, apiKey, search, perPage, page)
	if err != nil {
		return nil, err
	}
	return ResponseSearchPhoto(rawBody)
}

func SearchPhoto(apiKey, search string, perPage int, page int) (*Images, error) {
	return SearchPhotoWithHTTPClient(http.DefaultClient, apiKey, search, perPage, page)
}
