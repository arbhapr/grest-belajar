package app

import "grest.dev/grest"

// HttpClient creates and returns a new instance of httpClientUtil.
// It takes two parameters: method (HTTP method) and url (URL).
// The function initializes the Method and Url fields of the httpClientUtil instance and returns it.
func HttpClient(method, url string) *httpClientUtil {
	hc := &httpClientUtil{}
	hc.Method = method
	hc.Url = url
	return hc
}

// httpClientUtil represents a utility for making HTTP requests.
// It embeds the grest.HttpClient type, which provides additional functionality for making HTTP requests.
type httpClientUtil struct {
	grest.HttpClient
}
