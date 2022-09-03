package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func main() {
	inCSV, err := os.Open("data/input.csv")
    defer inCSV.Close()

    if err != nil {
        panic(err)
    }

    inLines, err := csv.NewReader(inCSV).ReadAll()
    if err != nil {
        panic(err)
    }

	outLines := [][]string{inLines[0]}

	const changesIndex = 4
	for i := 1; i < len(inLines); i++ {
		n, err := strconv.ParseInt(inLines[i][changesIndex], 10, 32)
		if err != nil {
			panic(err)
		}

		for chances := int64(0); chances < n; chances++ {
			outLines = append(outLines, inLines[i])
		}
	}

	outCSV, err := os.Create("data/output.csv")
    defer outCSV.Close()

    if err != nil {
        panic(err)
    }

	if err := csv.NewWriter(outCSV).WriteAll(outLines); err != nil {
		panic(err)
	}
}