package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	maxInt = 2147483647
)

func indexOfMin(dist []int, vis []bool) int {
	min, idx := maxInt, 0
	for i, v := range vis {
		if v == false && dist[i] < min {
			min = dist[i]
			idx = i
		}
	}
	if idx == 0 {
		return -1
	}
	return idx
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	apexCount, _ := strconv.Atoi(scn.Text())
	scn.Scan()
	from, _ := strconv.Atoi(scn.Text())
	scn.Scan()
	to, _ := strconv.Atoi(scn.Text())

	matrix := make([][]int, apexCount+1)
	for i := 1; i <= apexCount; i++ {
		matrix[i] = make([]int, apexCount+1)
		for j := 1; j <= apexCount; j++ {
			scn.Scan()
			value, _ := strconv.Atoi(scn.Text())
			matrix[i][j] = value
		}
	}
	inputFile.Close()

	visited := make([]bool, apexCount + 1)
	distance := make([]int, apexCount + 1)
	fromApex := make([]int, apexCount + 1)
	for i := 0; i < len(distance); i++ {
		distance[i] = maxInt
	}
	distance[from] = 0
	minDist := -1
	for true {
		idx := indexOfMin(distance, visited)
		if idx == -1 {
			break
		}
		visited[idx] = true
		for j := 1; j <= apexCount; j++ {
			if visited[j] == false && matrix[idx][j] != -1 && distance[j] > distance[idx] + matrix[idx][j] {
				distance[j] = distance[idx] + matrix[idx][j]
				fromApex[j] = idx
			}
		}
		if idx == to {
			minDist = distance[idx]
			break
		}
	}
	var ans []int
	if minDist == -1 {
		ans = []int{-1}
	}
	if minDist != -1 {
		idx := to
		for idx != 0 {
			ans = append([]int{idx}, ans...)
			idx = fromApex[idx]
		}
	}
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strings.Trim(fmt.Sprintf("%v", ans), "[]")))
	outputFile.Close()
}