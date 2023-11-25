package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var  (
	x uint64
	p uint64 = uint64(math.Pow10(9)) + 7
	h1, h2, xArr []uint64
)


func prepare(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		h1[i+1] = (h1[i]*x + uint64(arr[i]))%p
		h2[i+1] = (h2[i]*x + uint64(arr[size - i - 1]))%p
		xArr[i+1] = (xArr[i]*x)%p
	}
}

func isEqual(pos1, pos2, ln int) bool {
	return (
		(h1[pos1 + ln - 1] + h2[len(h2) - pos2 - ln] * xArr[ln])%p ==
		(h2[len(h2) - pos2] + h1[pos1 - 1]* xArr[ln])%p)
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	size, _ := strconv.Atoi(scn.Text())
	scn.Scan()
	colorCount, _ := strconv.Atoi(scn.Text())
	x = uint64(colorCount) + 1 
	array := make([]int, size)
	for i := 0; i < size; i++ {
		scn.Scan()
		value, _ := strconv.Atoi(scn.Text())
		array[i] = value
	}
	inputFile.Close()
	h1, h2 = make([]uint64, size + 1), make([]uint64, size + 1)
	xArr = make([]uint64, size + 1)
	xArr[0] = 1
	prepare(array)
	var ans []int
	for i := size/2; i >= 1; i--{
		if isEqual(1, 1+i, i) {
			ans = append(ans, size - i)
		}
	}
	ans = append(ans, size)
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Trim(fmt.Sprintf("%v", ans), "[]")))
	outputFile.Close()
}