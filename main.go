package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {
	if len(os.Args[1:]) < 3 {
		fmt.Println("need to provide site, concurrency and max pages arguments")
		fmt.Println(("format: $site $concurrency $maxPages"))
		os.Exit(1)
	} else if len(os.Args[1:]) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := os.Args[1:][0]
	concurrency, err := strconv.Atoi(os.Args[1:][1])
	if err != nil {
		fmt.Println("invalid concurrency value")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[1:][2])
	if err != nil {
		fmt.Println("invalid page pages value")
		os.Exit(1)
	}
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
		maxPages:           maxPages,
	}

	fmt.Printf("starting crawl of: %s\n", baseUrl)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseUrl)
	cfg.wg.Wait()

	printReport(cfg.pages, baseUrl)
}
