package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args[1:]) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := os.Args[1:][0]
	fmt.Printf("starting crawl of: %s\n", baseUrl)
	fmt.Println(getHTML(baseUrl))
}
