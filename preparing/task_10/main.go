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
	ans := make([]string, size)
	for i, _ := range ans {
		ans[i] = "No"
	}
	for i := 0; i < size; i++ {
		scn.Scan()
		n, _ := strconv.Atoi(scn.Text())
		scn.Scan()
		a, _ := strconv.Atoi(scn.Text())
		scn.Scan()
		b, _ := strconv.Atoi(scn.Text())
		if n%a <= n/a*(b-a) {
			ans[i] = "Yes"
		}
	}

	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Join(ans, "\n")))
	outputFile.Close()
	inputFile.Close()
}