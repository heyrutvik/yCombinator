package main

import (
	"flag"
	"fmt"     // format output string
	"log"     // print error on stderr
	"strings" // string manipulation functions

	"github.com/heyrutvik/yCombinator/search" // command line arguments
)

func main() {
	sFlag := flag.String("s", "golang", "search flag")
	flag.Parse()
	searchString := strings.Replace(*sFlag, " ", "%20", -1)

	fmt.Println("Wait...\n")

	items, err := search.Get(searchString)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
