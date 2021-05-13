package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

type record struct {
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	AdjClose string
}

type table []record

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	tableCsv, err := os.Open("table.csv")
	defer tableCsv.Close()
	if err != nil {
		log.Fatal("Cannot open table.csv", err)
	}

	records, err := csv.NewReader(tableCsv).ReadAll()
	if err != nil {
		log.Fatal("Cannot read data from table.csv", err)
	}

	var tbl table
	records = records[1:]
	for _, row := range records {
		i := record{
			Date:     row[0],
			Open:     row[1],
			High:     row[2],
			Low:      row[3],
			Close:    row[4],
			Volume:   row[5],
			AdjClose: row[6],
		}
		tbl = append(tbl, i)
	}

	indexHtml, err := os.Create("index.html")
	defer indexHtml.Close()
	if err != nil {
		log.Fatal("Cannot create index.html", err)
	}

	err = tpl.ExecuteTemplate(indexHtml, "tpl.gohtml", tbl)
	if err != nil {
		log.Fatal("Cannot execute tpl.gohtml", err)
	}
}
