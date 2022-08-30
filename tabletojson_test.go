package tabletojson

import (
	"fmt"
	"github.com/yigit433/tabletojson"
	"log"
	"net/http"
	"testing"
)

func TestTableToJson(t *testing.T) {
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
