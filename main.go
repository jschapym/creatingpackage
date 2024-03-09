package main

import (
	"fmt"

	trimmedmean "github.com/jschapym/creatingpackage"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := trimmedmean.TrimmedMean(data, 0.1)
	fmt.Println("Trimmed Mean:", result)
}
