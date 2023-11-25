package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func radixSort(arr []string, r int) ([]string, [][]string) {
	phase := make([][]string, 10)
	for _, v := range arr {
		phase[v[r] - 48] = append(phase[v[r] - 48], v)
	}
	var out []string
	for _, v := range phase {
		out = append(out, v...)
	}
	return out, phase
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	size, _ := strconv.Atoi(scn.Text())
	array := make([]string, size)
	for i:=0; i<size; i++ {
		scn.Scan()
		array[i] = scn.Text()
	}
	inputFile.Close()
	elemSize := len(array[0])
	var outArray []string
	outArray = append(outArray, "Initial array:", strings.Join(array, ", "), "**********")
	for i := elemSize - 1; i >= 0; i-- {
		outArray = append(outArray, fmt.Sprintf("Phase %d", elemSize - i))
		var phaseArray [][]string
		array, phaseArray = radixSort(array, i)
		for i, v := range phaseArray {
			value := strings.Join(v, ", ")
			if value == "" {
				value = "empty"
			}
			outArray = append(outArray, fmt.Sprintf("Bucket %d: ", i) + value)
		}
		outArray = append(outArray, "**********")
	}
	outArray = append(outArray, "Sorted array:", strings.Join(array, ", "))
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Join(outArray, "\n")))
	outputFile.Close()
}