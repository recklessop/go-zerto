package zerto

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	userAgent      = "go-zerto/" + libraryVersion
	defaultBaseURL = "https://206.55.216.101:9669/"
)

type (
	// Client for Zerto REST API
	Client struct {
		BaseURL   *url.URL
		UserAgent string
	
		httpClient *http.Client
	}

)
// PrintVersion of the library
func PrintVersion() {
	fmt.Println(userAgent)
}

