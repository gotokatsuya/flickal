package flickal

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type ImageURL struct {
	ThumbnailURL string `json:"thumbnail_url"`
	LargeURL     string `json:"large_url"`
}

type Images struct {
	Images []ImageURL `json:"images"`
	Pages  int        `json:"pages"`
	Page   int        `json:"page"`
}

func ResponseSearchPhoto(body []byte) (*Images, error) {
	bodyText := string(body)
	bodyText = strings.Replace(bodyText, "jsonFlickrApi(", "", -1)
	bodyText = strings.Replace(bodyText, ")", "", -1)
	decodedPhotoResponse, err := decodeSearchPhotoResponseBody([]byte(bodyText))
	if err != nil {
		return nil, err
	}
	return &Images{
		Images: buildImages(decodedPhotoResponse.Photos.Photo),
		Page:   decodedPhotoResponse.Photos.Page,
		Pages:  decodedPhotoResponse.Photos.Pages,
	}, nil
}

type Photo struct {
	ID       string `json:"id"`
	ServerID string `json:"server"`
	FarmID   int    `json:"farm"`
	Secret   string `json:"secret"`
}

type Photos struct {
	Photo []Photo `json:"photo"`
	Page  int     `json:"page"`
	Pages int     `json:"pages"`
}

type SearchPhotoResponse struct {
	Photos Photos `json:"photos"`
	Status string `json:"stat"`
}

func decodeSearchPhotoResponseBody(data []byte) (*SearchPhotoResponse, error) {
	response := &SearchPhotoResponse{}
	if err := json.Unmarshal(data, response); err != nil {
		return nil, err
	}
	if response.Status != "ok" {
		return nil, errors.New(response.Status)
	}
	return response, nil
}

const (
	imageURLFormat = "https://farm%d.staticflickr.com/%s/%s_%s_%s.jpg"
	thumbNailSize  = "t"
	largeSize      = "b"
)

func buildImages(photos []Photo) []ImageURL {
	var imageURLs []ImageURL
	for _, p := range photos {
		ThumbnailURL := fmt.Sprintf(imageURLFormat,
			p.FarmID, p.ServerID, p.ID, p.Secret, thumbNailSize)
		largeURL := fmt.Sprintf(imageURLFormat,
			p.FarmID, p.ServerID, p.ID, p.Secret, largeSize)
		imageURLs = append(imageURLs, ImageURL{
			ThumbnailURL: ThumbnailURL,
			LargeURL:     largeURL,
		})
	}
	return imageURLs
}
