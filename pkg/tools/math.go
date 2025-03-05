// math.go contains all integer mathematical operation not present in the
// standard math package.
package tools

import (
	"math"
	"math/rand"
	"time"
)

var (
	RandomRing *rand.Rand
)

func init() {
	InitRandom()
}

// -----------------------------------------------------------------------------
// Public functions
// -----------------------------------------------------------------------------

// Abs function returns the absolute value of an integer.
func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

// Sign function returns the positive or negative sign of an integer.
func Sign(a int) int {
	if a < 0 {
		return -1
	}
	return 1
}

func InitRandom() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	RandomRing = rand.New(source)
}

// Max function returns the maximum integer in a sequence of integers.
func Max(values ...int) int {
	max := math.MinInt
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// Min function returns the minimum integer in a sequence of integers.
func Min(values ...int) int {
	min := math.MaxInt
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// NilToInt functions returns a zero if the value is a nil.
func NilToInt(value any) int {
	if value == nil {
		return 0
	}
	return value.(int)
}

// SumSlice function returns the total sum for every entry in a slice or
// integer numbers.
func SumSlice(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
