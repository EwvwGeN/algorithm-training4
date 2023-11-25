package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func partition(array []int, left int, right int) (lowerEnd, highStart int) {
	if left == right {
		return left, left
	}
	lowerEnd = left
	highStart = right
	indx := left + rand.Intn(right - left)
	value := array[indx]
	for equalEnd := left; equalEnd < highStart; {
		switch{
		case array[equalEnd] == value:
			equalEnd++
		case array[equalEnd] < value:
			array[equalEnd], array[lowerEnd] = array[lowerEnd], array[equalEnd]
			lowerEnd++
			equalEnd++
		default:
			array[equalEnd], array[highStart - 1] = array[highStart - 1], array[equalEnd]
			highStart--
		}
	}
	return
}

func quickSort(arr []int, left int, right int) {
	if left == right {
		return
	}
	lwIndex, hgIndex := partition(arr, left, right)
	quickSort(arr, left, lwIndex)
	quickSort(arr, hgIndex, right)
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
	inputFile.Close()
	quickSort(array, 0, size)
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Trim(fmt.Sprintf("%v", array), "[]")))
	outputFile.Close()
}