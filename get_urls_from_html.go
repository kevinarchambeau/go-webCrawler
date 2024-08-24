package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return urls, fmt.Errorf("could not parse base url")
	}
	body := strings.NewReader(htmlBody)
	htmlNodes, err := html.Parse(body)
	if err != nil {
		return urls, fmt.Errorf("failed to get html nodes %v", err)
	}

	var links func(node *html.Node)
	links = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					parsedHref, err := url.Parse(a.Val)
					if err != nil {
						fmt.Println("could not parse as url")
						continue
					}
					resolvedURL := parsedBaseURL.ResolveReference(parsedHref)
					urls = append(urls, resolvedURL.String())
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			links(child)
		}
	}
	links(htmlNodes)
	return urls, nil
}
