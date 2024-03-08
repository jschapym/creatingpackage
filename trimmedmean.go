package trimmedmean

import (
	"sort"
)

func TrimmedMean(data []float64, trim float64) float64 {
	if len(data) == 0 {
		return 0
	}

	if trim < 0 || trim > 0.5 {
		return 0
	}

	sort.Float64s(data)

	n := len(data)
	trimSize := int(float64(n) * trim)

	if trimSize == 0 {
		return 0
	}

	data = data[trimSize : n-trimSize]

	sum := 0.0
	for _, v := range data {
		sum += v
	}

	return sum / float64(n-2*trimSize)
}
