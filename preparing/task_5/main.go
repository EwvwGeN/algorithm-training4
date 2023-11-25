package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	size, _ := strconv.Atoi(scn.Text())
	students := make([]int, size)
	for i:=0; i<size;i++ {
		scn.Scan()
		value, _ := strconv.Atoi(scn.Text())
		students[i] = value
	}
	inputFile.Close()
	diff := make([]int, size)
	sum := students[0]
	for i := 1; i < size; i++ {
		diff[i] = i*students[i] - sum
		sum += students[i]
	}
	sum = students[len(students) - 1]
	for i := len(students) - 2; i >= 0; i-- {
		diff[i] += sum - (len(students) - 1 - i)*students[i]
		sum += students[i]
	}

	ans := make([]string, size)
	for i := 0; i < size; i++ {
		ans[i] = strconv.Itoa(diff[i])
	}
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Join(ans, " ")))
	outputFile.Close()
}
