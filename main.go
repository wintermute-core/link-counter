package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/denis256/link-counter/core"
	"github.com/denis256/link-counter/links"
)

const DefaultWorkers = 2

func main() {
	urlsFile := flag.String("urls", "", "Text file with URLs to process.")
	output := flag.String("output", "", "Output file. Stdout by default.")
	jsonFormat := flag.Bool("json", false, "Enable JSON output")
	workers := flag.Int("workers", DefaultWorkers, "Number of parallel workers to process links")

	flag.Parse()

	if *urlsFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("Processing file %s", *urlsFile)
	list, err := links.ReadLinksFromFile(*urlsFile)

	if err != nil {
		panic(err)
	}

	if len(list) == 0 {
		log.Fatal("Empty URL list")
	}

	linkCounter := core.LinkCounter{Workers: *workers}
	results := linkCounter.Scan(list)

	var outputFile = os.Stdout

	if len(*output) != 0 {
		log.Printf("Saving output to %s", *output)
		outputFile, err = os.Create(*output)

		if err != nil {
			panic(err)
		}
		defer outputFile.Close()
	}

	if *jsonFormat {
		jsonOutput(results, outputFile)
	} else {
		humanOutput(results, outputFile)
	}
}

func humanOutput(results []core.ScanResult, outputFile *os.File) {
	for _, result := range results {
		_, err := fmt.Fprintf(outputFile,
			"page_url=%s internal_links_num=%d external_links_num=%d success=%v error_message=%s\n",
			result.PageURL, result.Internal, result.External, result.Success, result.Error)

		if err != nil {
			panic(err)
		}
	}
}

func jsonOutput(results []core.ScanResult, outputFile *os.File) {
	result, err := json.Marshal(results)

	if err != nil {
		panic(err)
	}

	_, err = outputFile.Write(result)

	if err != nil {
		panic(err)
	}
}
