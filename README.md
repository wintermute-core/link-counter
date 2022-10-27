# link-counter

CLI applications to count internal and external links on web pages.

Usage:
```
./link-counter 
Link counter
  -urls string
        Text file with URLs to process.
  -output string
        Output file. Stdout by default.
  -json
        Enable JSON output
  -workers int
        Number of parallel workers to process links (default 2)
```

Example input file:
```
https://universal-development.com/contacts/
https://universal-development.com/
http://universal-development.com/services/
http://universal-development.com/products/

```

Example executions:
```
./link-counter -urls urls.txt
2022/10/27 15:23:50 Processing file urls.txt
2022/10/27 15:23:50 Processing https://universal-development.com/contacts/
2022/10/27 15:23:50 Processing https://universal-development.com/
2022/10/27 15:23:50 Processing http://universal-development.com/services/
2022/10/27 15:23:50 Processing http://universal-development.com/products/
page_url=https://universal-development.com/contacts/ internal_links_num=5 external_links_num=7 success=true error_message=
page_url=https://universal-development.com/ internal_links_num=4 external_links_num=0 success=true error_message=
page_url=http://universal-development.com/services/ internal_links_num=5 external_links_num=1 success=true error_message=
page_url=http://universal-development.com/products/ internal_links_num=5 external_links_num=0 success=true error_message=


./link-counter -urls urls.txt -json -output file.json
2022/10/27 15:24:12 Processing file urls.txt
2022/10/27 15:24:12 Processing https://universal-development.com/contacts/
2022/10/27 15:24:12 Processing https://universal-development.com/
2022/10/27 15:24:12 Processing http://universal-development.com/services/
2022/10/27 15:24:12 Processing http://universal-development.com/products/
2022/10/27 15:24:12 Saving output to file.json
```
Example output in JSON format:
```
[
  {
    "page_url": "https://universal-development.com/contacts/",
    "internal_links_num": 5,
    "external_links_num": 7,
    "success": true,
    "error_message": ""
  },
  {
    "page_url": "https://universal-development.com/",
    "internal_links_num": 4,
    "external_links_num": 0,
    "success": true,
    "error_message": ""
  },
  {
    "page_url": "http://universal-development.com/services/",
    "internal_links_num": 5,
    "external_links_num": 1,
    "success": true,
    "error_message": ""
  },
  {
    "page_url": "http://universal-development.com/products/",
    "internal_links_num": 5,
    "external_links_num": 0,
    "success": true,
    "error_message": ""
  }
]
```

# Development

Dependencies:
  * go 1.19
  * make
  * goimports
  * golangci-lint

Installation of dependencies:
```
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

```

Make commands:
  * `make test`
  * `make fmt`
  * `make run-lint`
  * `make build`

Building application:
```
make build

# execute binary
./link-counter
```

## Application structure

Packages:
  * `page` - downloading and parsing of HTML pages
  * `links` - processing of links, identification of links which points on external and internal resources
  * `core` - logic for parallel processing of links


## Future work

Future improvements:
  * configurable HTTP client
  * http redirects follow of HTTP redirects 
  * usage of HTTP proxies and retries in case of errors
  * processing of javascript on pages
  * separation of unit and integration tests
  * more flexible configuration for handling of urls and links
  * exposing of runtime metrics
  * more output formats
  * CICD integration and packaging of release binaries

