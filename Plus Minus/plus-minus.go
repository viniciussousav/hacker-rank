package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	positives, negatives, zeros := 0, 0, 0

	for _, n := range arr {
		if n > 0 {
			positives++
			continue
		}
		if n < 0 {
			negatives++
			continue
		}

		zeros++
	}

	arrSize := float64(len(arr))

	positiveRatio := float64(positives) / arrSize
	negativeRatio := float64(negatives) / arrSize
	zeroRatio := float64(zeros) / arrSize

	fmt.Printf("%.6f\n", positiveRatio)
	fmt.Printf("%.6f\n", negativeRatio)
	fmt.Printf("%.6f", zeroRatio)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
