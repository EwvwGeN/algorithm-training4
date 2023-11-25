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
	buf := make([]byte, 1024*1024)
	scn.Buffer(buf, 1024*1024)
	scn.Scan()
	str := scn.Text()
	inputFile.Close()
	str = "#"+strings.Join(strings.Split(str, ""), "#")+"#"
	size := len(str)
	ans := make([]int, size)
  	l, r := 0,0
	count := 0
	for i := 1; i < size; i++ {
        if (i < r) {
			ans[i] = min(r - i + 1, ans[l + r - i]);
		}
        for (i - ans[i] >= 0 && i + ans[i] < size && str[i - ans[i]] == str[i + ans[i]]) {
			ans[i]++
		}
    	if (i + ans[i] - 1 > r) {
			l = i - ans[i] + 1
			r = i + ans[i] - 1
		}
        count+=ans[i]/2
    }
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strconv.Itoa(count)))
	outputFile.Close()
}