package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
)

func solve1() {
	csvFile, _ := os.Open("historicalPrices.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var stockPrices []Stock
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		date, _ := time.Parse("1/2/06", line[0])
		stockPrices = append(stockPrices, Stock{
			Date:       date,
			OpenPrice:  line[1],
			HighPrice:  line[2],
			LowPrice:   line[3],
			ClosePrice: line[4],
		})
	}

	// for _, stock := range stockPrices {
	// 	printVal(stock.HighPrice)
	// }
	// startingGeld := 10000
	//Starting on some date, we buy a certain percentage of stocks at some certain interval, and see the value of them some years later
	// today
	printVal(getPriceNearDate(getDateForString("12/20/13	"), stockPrices))
}

func getDateForString(dateString string) time.Time {
	date, _ := time.Parse("1/2/06", dateString)
	return date
}

func getPriceNearDate(date time.Time, stockPrices []Stock) string {
	for _, stock := range stockPrices {
		if date.After(stock.Date) {
			continue
		}
		return stock.OpenPrice
	}
	return stockPrices[len(stockPrices)-1].OpenPrice
}

//Date, Open, High, Low, Close
type Stock struct {
	Date       time.Time `json:"date"`
	OpenPrice  string    `json:"openprice"`
	HighPrice  string    `json:"highprice"`
	LowPrice   string    `json:"lowprice"`
	ClosePrice string    `json:"closeprice"`
}
