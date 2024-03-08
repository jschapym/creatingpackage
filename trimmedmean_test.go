package trimmedmean_test

import (
	"sort"
	"testing"
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

func TestTrimmedMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	trim := 0.2
	want := 3.0
	got := TrimmedMean(data, trim)
	if got != want {
		t.Errorf("TrimmedMean(%v, %v) = %v; want %v", data, trim, got, want)
	}
}
