package flickal

import (
	"flag"
	"os"
	"testing"
)

var apiKey = flag.String("apikey", os.Getenv("flicker_apikey"), "")

func TestSearchPhoto(t *testing.T) {
	res, err := SearchPhoto(*apiKey, "bird", 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
