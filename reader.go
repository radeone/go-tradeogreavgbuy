package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
//    "unsafe"
)

type orderLine struct {
    side bool
    pair string
    id float64
    size float64
    price float64
}

type myBook struct {
   pairs map[string] []orderLine
   weights []float64
   sizes []float64
}

func isBuy(order string) bool {
	if order == "BUY" {
		return true
	} else {
		return false
	}
}
func dumpFloat(order string) float64 {
	dump, _ := strconv.ParseFloat(order,64)
	return dump
}

func readCSV(filename string) []orderLine { //returns all orderlines

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
    books := make([]orderLine,0)
    for _, line := range lines {
        data := orderLine{
            side: isBuy(line[0]),
            pair: line[1],
            id: dumpFloat(line[2]),
	    size: dumpFloat(line[3]),
	    price: dumpFloat(line[4]),
        }
	books = append(books,data)
   }
   return books
}

func sortBooks(boox []orderLine) map[string][]orderLine {
	bythepair := make(map[string] []orderLine)
	for _, line := range boox {
		if line.side { // only adds buys to database sorted by pair in map
			bythepair[line.pair] = append(bythepair[line.pair],line)
		}
	}
	return bythepair
}

func avgWeight(pair []orderLine) float64 {
	var totalSize float64
	cpo := make([]float64,0)
	var avg float64
	for _, line := range pair { //calculates all the buys
		totalSize += line.size
	}
	for _, line := range pair { //calculates all convexes and the sum of convex * price of order
		cx := line.size/totalSize
		cpo = append(cpo,cx)
		avg += cx*line.price
	}
	return avg*1e8 //satoshi math
}

func main() {
	unsorted := readCSV("trades.csv")
	sorted := sortBooks(unsorted)
	fmt.Println(avgWeight(sorted["BTC-XHV"]))
//	debug shit
//	fmt.Println(len(sorted))
//	fmt.Println(unsafe.Sizeof(shit)) seeings sizes lol

}
