package tabletojson

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

func Convert(reader io.ReadCloser) ([]Table, error) {
	tables := []Table{}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return tables, err
	}

	doc.Find("table").Each(func(_ int, tableSelection *goquery.Selection) {
		table := Table{}

		tableSelection.Find("tr").Each(func(_ int, trSelection *goquery.Selection) {
			row := Row{}

			trSelection.Find("th").Each(func(thi int, thSelection *goquery.Selection) {
				row = append(row, Column{
					ParentName: strings.TrimSpace(thSelection.Text()),
				})
			})

			trSelection.Find("td").Each(func(tdi int, tdSelection *goquery.Selection) {
				if len(row) == 0 || len(row) == tdi {
					row = append(row, Column{
						ParentValue: strings.TrimSpace(tdSelection.Text()),
					})
				} else {
					row[tdi].ParentValue = strings.TrimSpace(tdSelection.Text())
				}
			})

			table = append(table, row)
		})

		tables = append(tables, table)
	})

	return tables, nil
}
