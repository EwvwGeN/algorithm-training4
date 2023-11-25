package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
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
	scn.Scan()
	h = make([]int, len(str) + 1)
	xArr = make([]int, len(str) + 1)
	xArr[0] = 1
	prepare(str)
	count, _ := strconv.Atoi(scn.Text())
	ans := make([]string, count)
	for i:=0; i<count; i++ {
		scn.Scan()
		action := strings.Split(scn.Text(), " ")
		len, _ := strconv.Atoi(action[0])
		pos1,_ := strconv.Atoi(action[1])
		pos2,_ := strconv.Atoi(action[2])
		ans[i] = "no"
		if isEqual(pos1 + 1, pos2 + 1, len) {
			ans[i] = "yes"
		}
		
	}
	inputFile.Close()

	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Join(ans, "\n")))
	outputFile.Close()
}