package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %v\n", baseURL)
	fmt.Println("=============================")

	keys := make([]int, 0, len(pages))
	sortedList := map[int][]string{}

	for key, value := range pages {
		sortedList[value] = append(sortedList[value], key)
	}

	for key := range sortedList {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for i := len(keys) - 1; i >= 0; i-- {
		sort.Strings(sortedList[keys[i]])
		for _, value := range sortedList[keys[i]] {
			fmt.Printf("Found %v internal links to %v\n", keys[i], value)
		}
	}

}
