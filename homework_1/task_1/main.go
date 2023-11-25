package main

import (
	"bufio"
	"os"
	"strconv"
)

func partition(p int, array []int, left int, right int) int {
	lowerEnd, equalEnd := left, left
	for i := left; i<right; i++ {
		if array[i] > p {continue}
		if array[i] == p {
			array[i], array[equalEnd] = array[equalEnd], array[i]
			equalEnd++
			continue
		}
		if array[i] < p {
			array[i], array[lowerEnd] = array[lowerEnd], array[i]
			if lowerEnd == equalEnd {
				lowerEnd++
				equalEnd++
				continue
			}
			lowerEnd++
			array[i], array[equalEnd] = array[equalEnd], array[i]
			equalEnd++
			continue
		}
	}
	if lowerEnd == 0 {
		return lowerEnd
	}
	return lowerEnd
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	size, _ := strconv.Atoi(scn.Text())
	array := make([]int, size)
	for i:=0; i<size; i++ {
		scn.Scan()
		value, _ := strconv.Atoi(scn.Text())
		array[i] = value
	}
	scn.Scan()
	p, _ := strconv.Atoi(scn.Text())
	inputFile.Close()
	lowerEndIndex := partition(p, array, 0, size)
	coutOther := strconv.Itoa(size - lowerEndIndex)
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strconv.Itoa(lowerEndIndex) + "\n" + coutOther))
	outputFile.Close()
}