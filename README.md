## Table To Json [![Go Report Card](https://goreportcard.com/badge/github.com/yigit433/tabletojson-go)](https://goreportcard.com/report/github.com/yigit433/tabletojson-go)
Attempts to convert HTML tables into JSON. 
### Example
```go
package main

import (
    "fmt"
    "github.com/yigit433/tabletojson-go"
    "net/http"
)

func main() {
    res, err := http.Get("https://www.doviz.com/kripto-paralar")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

    tables, err := tabletojson.Convert(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(tables)
}
```
