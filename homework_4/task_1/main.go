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
	inputLen, _ := strconv.Atoi(scn.Text())
	ans := getPermutation(inputLen)
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Join(ans, "\n")))
	outputFile.Close()
}

func findMaxIndex(permutation string) int {
	for i := len(permutation) - 2; i > -1; i-- {
		if permutation[i] < permutation[i+1] {
			return i
		}
	}
    return -1
}

func findBiggerIndex(permutation string, element byte) int {
	for i := len(permutation) - 1; i > -1; i-- {
        if permutation[i] > element {
			return i
		}  
	}
    return -1
}

func swap(permutation string, i, j int) string {
	buf := []byte(permutation)
	buf[i], buf[j] = buf[j], buf[i]
	return string(buf)
}

func reverseStr(permutation string, index int) string {
	n := len(permutation)
	rArr := []byte(permutation)
	for i := 0; i < n/2; i++ { 
		rArr[i], rArr[n-1-i] = rArr[n-1-i], rArr[i] 
	
	}
	buf := string(rArr)
	permutation = permutation[:index+1] + buf[:n-index - 1]
        
    return permutation
}

func getPermutation(n int) (out []string) {
	var str string
	for i := 1; i <= n; i++ {
		str = str + strconv.Itoa(i)
	}
    out = append(out, str)
    index := findMaxIndex(str)
    for index != -1 {
		element := str[index]
        swap_index := findBiggerIndex(str, element)
        str = swap(str, index, swap_index)
        str = reverseStr(str, index)
        out = append(out, str)
        index = findMaxIndex(str)
	}
    return
}
    


