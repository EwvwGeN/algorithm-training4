package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	a, _ := strconv.Atoi(scn.Text())
	scn.Scan()
	b, _ := strconv.Atoi(scn.Text())
	scn.Scan()
	n, _ := strconv.Atoi(scn.Text())
	inputFile.Close()
	bcount := b/n
	if b % n > 0 {
		bcount++
	}
 	ans := "No"
	if a > bcount {
		ans = "Yes"
	}
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(ans))
	outputFile.Close()
	inputFile.Close()
}