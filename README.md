# link-counter

Example output in text format:
```
page_url=https://universal-development.com/ internal_links_num=4 external_links_num=0 success=true error_message=
page_url=https://universal-development.com/contacts/ internal_links_num=5 external_links_num=7 success=true error_message=
page_url=http://universal-development.com/services/ internal_links_num=5 external_links_num=1 success=true error_message=
page_url=http://universal-development.com/products/ internal_links_num=5 external_links_num=0 success=true error_message=

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
  * golangci-lint

Make commands:
  * `make test`
  * `make fmt`
  * `make run-lint`
  * `make build`

Execute final binary: 
```
./link-counter
```