package main

import (
	"fmt"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) map[string]int {
	baseURL, err := normalizeURL(rawBaseURL)
	if err != nil {
		return pages
	}
	currentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return pages
	}

	if !strings.Contains(currentURL, baseURL) {
		return pages
	}

	if _, ok := pages[currentURL]; ok {
		pages[currentURL]++
		return pages
	} else {
		pages[currentURL] = 1
	}
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("request failed: %v\n", err)
		return pages
	}

	list, err := getURLsFromHTML(body, rawBaseURL)
	if err != nil {
		return pages
	}
	for _, item := range list {
		fmt.Printf("item is: %v\n", item)
		crawlPage(rawBaseURL, item, pages)
	}

	return pages
}
