package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
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

	bits := strconv.FormatInt(n, 2)

	flippedBitsArr := make([]string, 32)

	for i := range flippedBitsArr {

		j := len(flippedBitsArr) - i - 1
		if j < len(bits) && bits[j] == '1' {
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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := flippingBits(n)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
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
