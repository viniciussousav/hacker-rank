package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
 * Complete the 'flippingBits' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER n as parameter.
 */

func flippingBits(n int64) int64 {

	bits := strconv.FormatUint(uint64(n), 2)
	fmt.Println("[bits]:", bits)

	flippedBitsArr := make([]string, 32)

	for i := range flippedBitsArr {

		j := len(flippedBitsArr) - i - 1
		if j < len(bits) && bits[len(bits)-j-1] == '1' {
			flippedBitsArr[i] = "0"
			continue
		}

		flippedBitsArr[i] = "1"
	}

	flippedBitsString := strings.Join(flippedBitsArr, "")
	flippedBitsInt, _ := strconv.ParseUint(flippedBitsString, 2, 32)

	return int64(math.Abs(float64(flippedBitsInt)))
}

func main() {
	result := flippingBits(2147483647)
	fmt.Println("[Resultado]:", result)
}
