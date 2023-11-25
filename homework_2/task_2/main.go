package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var  (
	x int = 27
	p int = int(math.Pow10(9)) + 7
	h, xArr []int
)


func prepare(str string) {
	for i, v := range str {
		h[i+1] = (h[i]*x + int(v) - 97)%p
		xArr[i+1] = (xArr[i]*x)%p
	}
}

func isEqual(pos1, pos2, len int) bool {
	return (
		(h[pos1 + len - 1] + h[pos2 - 1] * xArr[len])%p ==
		(h[pos2 + len - 1] + h[pos1 - 1] * xArr[len])%p)
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	buf := make([]byte, 1024*1024)
	scn.Buffer(buf, 1024*1024)
	scn.Scan()
	str := scn.Text()
	inputFile.Close()
	size := len(str)
	h = make([]int, size + 1)
	xArr = make([]int, size + 1)
	xArr[0] = 1
	prepare(str)
	ans := size
	for i := 1; i<=size; i++{
		trigger := true
		for j:=1; j<=size + 1 - i; j+=i {
			if !isEqual(1, j, i) {
				trigger = false
				break
			}
				
		}
		if !isEqual(1, i * (size / i) + 1, size % i) {
			trigger = false
		}
		if trigger {
			ans = i
			break
		}
	}
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(fmt.Sprintf("%d", ans)))
	outputFile.Close()
}