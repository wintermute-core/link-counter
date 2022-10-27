package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const DefaultWorkers = 2

func main() {
	fmt.Println("Link counter")

	linksFile := flag.String("links", "", "Text file with URLs to process.")

	workers := flag.Int("workers", DefaultWorkers, "Number of parallel workers to process links")

	flag.Parse()

	if *linksFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("Processing file %s", *linksFile)

	fmt.Printf("linksFile: %s \n", *linksFile)
	fmt.Printf("workers: %d \n", *workers)
}
