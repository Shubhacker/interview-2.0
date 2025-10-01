package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type urlsInterface interface {
	Shortner(string) string
	ExpandUrl(string) string
}
type UrlStruct struct {
	ActualURL  map[string]string
	DynamicUrl string
}

func main() {
	u := UrlStruct{
		ActualURL:  make(map[string]string, 0),
		DynamicUrl: "test.ly",
	}
	shortUrl := u.Shortner("https://www.geeksforgeeks.org/go-language/generate-uuid-in-golang/")
	fmt.Println("Short URL --->", shortUrl)
	// Short URL ---> test.ly/73bcc1f4
	actualUrl := u.ExpandUrl(shortUrl)
	fmt.Println("Actual URL --->", actualUrl)
	// Actual URL ---> https://www.geeksforgeeks.org/go-language/generate-uuid-in-golang/
}

// Shortner, Will create new url and store current url in memory.
func (u UrlStruct) Shortner(actualUrl string) string {
	id := uuid.New()
	uuid := strings.Split(id.String(), "-")
	k := u.generateShortUrl(uuid[0])
	u.ActualURL[uuid[0]] = actualUrl
	return k
}

// Will generate new url with provided custom url + uuid (unique id)
func (u UrlStruct) generateShortUrl(id string) string {
	var shortUrl string
	shortUrl = u.DynamicUrl + "/" + id
	return shortUrl
}

// Fetch actual url from map
func (u UrlStruct) ExpandUrl(shortUrl string) string {
	id := strings.Split(shortUrl, "/")
	actualUrl, ok := u.ActualURL[id[1]]
	if !ok {
		return ""
	}

	return actualUrl
}
