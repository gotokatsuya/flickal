package flickal

func SearchPhoto(apiKey, search string, perPage int, page int) (*Images, error) {
	rawBody, err := RequestSearchPhoto(apiKey, search, perPage, page)
	if err != nil {
		return nil, err
	}
	return ResponseSearchPhoto(rawBody)
}
