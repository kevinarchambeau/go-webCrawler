package main

import (
	"fmt"
	"strings"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	// defer needs a function
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	if !strings.Contains(currentURL, cfg.baseURL.Host) {
		return
	}

	if !cfg.addPageVisit(currentURL) {
		return
	}

	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("request failed: %v\n", err)
		return
	}

	list, err := getURLsFromHTML(body, rawCurrentURL)
	if err != nil {
		return
	}
	for _, item := range list {
		cfg.wg.Add(1)
		go cfg.crawlPage(item)
	}

}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL]++
		return false
	} else {
		cfg.pages[normalizedURL] = 1
	}
	return true
}
