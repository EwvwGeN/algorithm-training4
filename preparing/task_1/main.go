package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func innerLog(l int) int {
	if l == 1 {return 0}
	return innerLog(l/2) + 1
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)

	scn.Scan()
	size, _ := strconv.Atoi(scn.Text())
	array := make([]int, size)

	scn.Scan()
	coutOperation, _ := strconv.Atoi(scn.Text())

	for i := 0; i < size; i++ {
		scn.Scan()
		value, _ := strconv.Atoi(scn.Text())
		array[i] = value
	}
	
	logs := make([]int, size)

	for i:=1; i<=size; i++ {
		logs[i-1] = innerLog(i)
	}

	sepTableMin := make([][]int, logs[size-1] + 1)
	sepTableMin[0] = array
	sepTableMax := make([][]int, logs[size-1] + 1)
	sepTableMax[0] = array
	outed := 0
	for i := 1; i < len(sepTableMin); i++ {
		outed += int(math.Pow(2, float64(i-1)))
		sepTableMin[i] = make([]int, size-outed)
		sepTableMax[i] = make([]int, size-outed)
		for j := 0; j < len(sepTableMin[i]); j++ {
			sepTableMin[i][j] = min(sepTableMin[i-1][j], sepTableMin[i-1][j+int(math.Pow(2, float64(i-1)))])
			sepTableMax[i][j] = max(sepTableMax[i-1][j], sepTableMax[i-1][j+int(math.Pow(2, float64(i-1)))])
		}
	}

	answers := make([]string, coutOperation)

	for i := 0; i<coutOperation; i++ {
		scn.Scan()
		left, _ := strconv.Atoi(scn.Text())
		scn.Scan()
		right, _ := strconv.Atoi(scn.Text())

		k := logs[right - left]
		minInRange := min(sepTableMin[k][left], sepTableMin[k][right - int(math.Pow(2, float64(k))) + 1])
		maxInRange := max(sepTableMax[k][left], sepTableMax[k][right - int(math.Pow(2, float64(k))) + 1])
		if maxInRange == minInRange {
			answers[i] = "NOT FOUND"
			continue
		}
		answers[i] = strconv.Itoa(maxInRange)
	}

	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Join(answers, "\n")))
	outputFile.Close()
	inputFile.Close()
}