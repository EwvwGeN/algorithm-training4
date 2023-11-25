package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	buf := make([]byte, 1024*1024)
	scn.Buffer(buf, 1024*1024)
	scn.Scan()
	str := scn.Text()
	inputFile.Close()
	size := len(str)
	zArr := make([]int, size)
	lPos, rPos := 0, 0
	for i:= 1; i<size; i++ {
		zArr[i] = max(0, min(zArr[i - lPos], rPos - i))
		for i + zArr[i] < size && str[zArr[i]] == str[i + zArr[i]]{
			zArr[i]++
		}
		if i + zArr[i] > rPos {
			lPos, rPos = i, i + zArr[i]
		}
	}

	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Trim(fmt.Sprintf("%v", zArr), "[]")))
	outputFile.Close()
}