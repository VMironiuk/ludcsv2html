package main

import (
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
	tbl := table{
		record{
			Date:     "2015-07-09",
			Open:     "523.119995",
			High:     "523.77002",
			Low:      "520.349976",
			Close:    "520.679993",
			Volume:   "1839400",
			AdjClose: "520.679993",
		},
		record{
			Date:     "2015-07-08",
			Open:     "521.049988",
			High:     "522.734009",
			Low:      "516.109985",
			Close:    "516.830017",
			Volume:   "1264600",
			AdjClose: "516.830017",
		},
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
