package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(arr1, arr2, initArr []int) {
	firstPos, secondPos := 0, 0
	firstLen, secondLen := len(arr1), len(arr2)
	for currentPos := 0 ; currentPos < len(initArr); currentPos++{
		if firstPos == firstLen {
			initArr[currentPos] = arr2[secondPos]
			secondPos++
			continue
		}
		if secondPos == secondLen {
			initArr[currentPos] = arr1[firstPos]
			firstPos++
			continue
		}
		if	arr2[secondPos] < arr1[firstPos] {
			initArr[currentPos] = arr2[secondPos]
			secondPos++
			continue
		}
		initArr[currentPos] = arr1[firstPos]
		firstPos++
	}
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	firstSize, _ := strconv.Atoi(scn.Text())
	firstArray := make([]int, firstSize)
	for i:=0; i<firstSize; i++ {
		scn.Scan()
		value, _ := strconv.Atoi(scn.Text())
		firstArray[i] = value
	}
	scn.Scan()
	secondSize, _ := strconv.Atoi(scn.Text())
	secondArray := make([]int, secondSize)
	for i:=0; i<secondSize; i++ {
		scn.Scan()
		value, _ := strconv.Atoi(scn.Text())
		secondArray[i] = value
	}
	inputFile.Close()
	arr := make([]int, firstSize + secondSize)
	merge(firstArray, secondArray, arr)
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Trim(fmt.Sprintf("%v", arr), "[]")))
	outputFile.Close()
}