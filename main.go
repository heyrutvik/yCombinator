package main

import (
	"flag" // command line arguments
	"fmt"  // format output string
	"github.com/heyrutvik/yCombinator/search"
	"log"     // print error on stderr
	"strings" // string manipulation functions
)

func main() {
	sFlag := flag.String("s", "golang", "search flag")
	flag.Parse()
	search := strings.Replace(*sFlag, " ", "%20", -1)

	fmt.Println("Wait...\n")

	items, err := hn.Get(search)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
