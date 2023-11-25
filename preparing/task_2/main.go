package main

import (
	"bufio"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	for b!=0 {
		tmp := a
		a = b
		b = tmp % a
	}
	return a
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	array := make([]int, 4)
	for i := 0; i < 4; i++ {
		scn.Scan()
		value, _ := strconv.Atoi(scn.Text())
		array[i] = value
	}
	inputFile.Close()

	fraction := make([]int, 2)
	fraction[0] = array[0]*array[3] + array[2]*array[1]
	fraction[1] = array[1]*array[3]

	divisor := gcd(fraction[0], fraction[1])
	fraction[0] = fraction[0]/divisor
	fraction[1] = fraction[1]/divisor
	ans := strconv.Itoa(fraction[0]) + " " + strconv.Itoa(fraction[1])
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(ans))
	outputFile.Close()
}