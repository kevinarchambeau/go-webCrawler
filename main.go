package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args[1:]) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := os.Args[1:][0]
	concurrency := 1
	parsedBaseUrl, err := url.Parse(baseUrl)
	if err != nil {
		fmt.Println("unable to parse url")
		os.Exit(1)
	}
	cfg := config{
		pages:              map[string]int{},
		baseURL:            parsedBaseUrl,
		concurrencyControl: make(chan struct{}, concurrency),
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
	}

	fmt.Printf("starting crawl of: %s\n", baseUrl)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseUrl)
	cfg.wg.Wait()

	fmt.Println(cfg.pages)
}
