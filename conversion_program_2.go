package main

import (
	"fmt"
	"math"
)

func main() {
	result := IntFromFloat64(9.4)
	fmt.Println(result)
}

func IntFromFloat64(floatValue float64) int {
	if math.MinInt32 <= floatValue && floatValue <= math.MaxInt32 {
		whole, fraction := math.Modf(floatValue)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}

	panic(fmt.Sprintf("%g is out of the int32 range", floatValue))
}
