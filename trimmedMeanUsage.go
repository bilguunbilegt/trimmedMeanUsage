package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bilguunbilegt/trimmedMean"
)

func readNumbersFromCSV(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV: %w", err)
	}

	var numbers []float64
	for _, record := range records {
		for _, value := range record {
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("could not parse number %s: %w", value, err)
			}
			numbers = append(numbers, num)
		}
	}

	return numbers, nil
}

func main() {
	filePathInt := "sample_int.csv" // Path to the sample CSV file

	// Read numbers from the CSV file
	numbers, err := readNumbersFromCSV(filePathInt)
	if err != nil {
		log.Fatalf("Error reading numbers from CSV: %v", err)
	}

	// Calculate the trimmed mean with a trim percentage of 5 (5%)
	mean, err := trimmedMean.TrimmedMean(numbers, 5, 5)
	if err != nil {
		fmt.Println("Error calculating trimmed mean for ints:", err)
	} else {
		fmt.Println("Trimmed Mean of integers:", mean)
	}

	filePathFloat := "sample_float.csv" // Path to the sample CSV file

	// Read numbers from the CSV file
	floats, err := readNumbersFromCSV(filePathFloat)
	if err != nil {
		log.Fatalf("Error reading numbers from CSV: %v", err)
	}

	// Calculate the trimmed mean with a trim percentage of 5 (5%)
	mean_float, err := trimmedMean.TrimmedMean(floats, 5, 5)
	if err != nil {
		fmt.Println("Error calculating trimmed mean for floats:", err)
	} else {
		fmt.Println("Trimmed Mean of floats:", mean_float)
	}
}
