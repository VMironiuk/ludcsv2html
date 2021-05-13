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
	writeRecords("index.html", "tpl.gohtml", readRecords("table.csv"))
}

func readRecords(filename string) table {
	tableCsv, err := os.Open(filename)
	defer tableCsv.Close()
	if err != nil {
		log.Fatal("Cannot open csv file", err)
	}

	records, err := csv.NewReader(tableCsv).ReadAll()
	if err != nil {
		log.Fatal("Cannot read data from csv file", err)
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

	return tbl
}

func writeRecords(filename string, template string, tbl table) {
	indexHtml, err := os.Create(filename)
	defer indexHtml.Close()
	if err != nil {
		log.Fatal("Cannot create html file", err)
	}

	err = tpl.ExecuteTemplate(indexHtml, template, tbl)
	if err != nil {
		log.Fatal("Cannot execute gohtml file", err)
	}
}
