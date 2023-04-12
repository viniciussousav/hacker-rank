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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {

	period := s[len(s)-2:]
	times := strings.Split(s[:len(s)-2], ":")

	hour, _ := strconv.Atoi(times[0])

	if period == "AM" && hour == 12 {
		return strings.Join([]string{"00", times[1], times[2]}, ":")
	}

	if (period == "PM" && hour == 12) || period == "AM" {
		return s[:len(s)-2]
	}

	militaryHour := strconv.Itoa(hour + 12)
	return strings.Join([]string{militaryHour, times[1], times[2]}, ":")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
