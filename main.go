package main

import "fmt"

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

	fmt.Println(tbl)
}
