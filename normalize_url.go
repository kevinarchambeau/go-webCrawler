package main

import "net/url"

func normalizeURL(unparsedURL string) (string, error) {
	parsedUrl, err := url.Parse(unparsedURL)
	if err != nil {
		return "", err
	}
	if parsedUrl.Path[len(parsedUrl.Path)-1:] == "/" {
		return parsedUrl.Host + parsedUrl.Path[0:len(parsedUrl.Path)-1], nil
	}

	return parsedUrl.Host + parsedUrl.Path, nil
}
