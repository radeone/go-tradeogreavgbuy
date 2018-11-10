package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

type CsvLine struct {
    side string
    pair string
    id string
    size string
    paid string
}

func main() {

    filename := "trades.csv"

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        panic(err)
    }

    // Loop through lines & turn into object
    var buytotal float64
    var selltotal float64
    buyz := make([]float64,0) //buy sizes
    sellz := make([]float64,0) //sell sizes
    bPrice := make([]float64,0) //buy prices
    sPrice := make([]float64,0) //sell prices
    countbuyz := 0
    for _, line := range lines {
        data := CsvLine{
            side: line[0],
            pair: line[1],
            id: line[2],
	    size: line[3],
	    paid: line[4],
        }

	//choose a pair from whatevs
	if data.pair == "BTC-RVN" {
		size, _ := strconv.ParseFloat(data.size,64)
		paid, _ := strconv.ParseFloat(data.paid,64)
		if data.side == "BUY" {
			buytotal += size
			buyz = append(buyz,size)
			bPrice = append(bPrice,paid)
			countbuyz++
		} else {
			selltotal += size
			sellz = append(sellz,size)
			sPrice = append(sPrice,paid)
		}
    	}
    }
    fmt.Println("Total bought: ",buytotal,"\ntotal buys: ",countbuyz)
    fmt.Println("Total sold: ",selltotal)
    fmt.Println("Difference: ",buytotal-selltotal)

	cons := make([]float64,0)
	for i := 0; i<countbuyz;i++ {
		cons = append(cons,buyz[i]/buytotal) //calculate and create convex slice size of order divided by total buy order size
	}
	var avgprice float64
	for i := 0; i<countbuyz;i++ {
		avgprice += cons[i]*bPrice[i] //sum of convex of order multiplied by the price of order
	}
	fmt.Println("avg price (sats): ",avgprice*1e8) //satoshi math 1e8
}
