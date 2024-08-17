package main

import (
	"fmt"

	"github.com/bilguunbilegt/trimmedMean"
)

func main() {
	numbers := []float64{1, 2, 3, 4, 5, 6, 7}
	mean, err := trimmedMean.TrimmedMean(numbers, 1, 1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Trimmed Mean:", mean)
	}
}
