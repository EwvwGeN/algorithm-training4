package main

import (
	"bufio"
	"os"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	buf := make([]byte, 0, 64*1024)
	scn.Buffer(buf, 1024*1024)
	words := make([]string, 2)
	for i:=0; i<2;i++ {
		scn.Scan()
		value := scn.Text()
		words[i] = value
	}
	inputFile.Close()
	ans := "NO"

	usedLetters := make(map[byte]int, 24)

	if len(words[0]) == len(words[1]) {
		ans = "YES"
		for i := range words[0] {
			usedLetters[words[0][i]] += 1
			usedLetters[words[1][i]] -= 1
		}
		for _, v := range usedLetters {
			if v != 0 {
				ans = "NO"
				break
			}
		}
	}

	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(ans))
	outputFile.Close()
}
