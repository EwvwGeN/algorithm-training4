package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	a1 := mergeSort(arr[0:len(arr)/2])
	a2 := mergeSort(arr[len(arr)/2:])
	return merge(a1, a2)
}

func merge(arr1, arr2 []int) (out []int) {
	firstPos, secondPos := 0, 0
	firstLen, secondLen := len(arr1), len(arr2)
	out = make([]int, firstLen+ secondLen)
	for currentPos := 0 ; currentPos < len(out); currentPos++{
		if firstPos == firstLen {
			out[currentPos] = arr2[secondPos]
			secondPos++
			continue
		}
		if secondPos == secondLen {
			out[currentPos] = arr1[firstPos]
			firstPos++
			continue
		}
		if	arr2[secondPos] < arr1[firstPos] {
			out[currentPos] = arr2[secondPos]
			secondPos++
			continue
		}
		out[currentPos] = arr1[firstPos]
		firstPos++
	}
	return
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
	out := mergeSort(array)
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Trim(fmt.Sprintf("%v", out), "[]")))
	outputFile.Close()
}