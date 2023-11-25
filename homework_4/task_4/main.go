package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

type heapStruct []int

func (h heapStruct) Add(apex int, dist int) {
	h[0] = h[0] + dist
	h[1] -= matrix[apex][0]
	h[2] = apex
	h[len(h)-2]--
	if h[len(h)-2] == 0 {
		h[len(h)-1] = 1
	}
	h[2+apex] = 1
}

func (h heapStruct) CheckVis(apex int) bool {
	return h[2+apex] == 1
}

type heapInit []heapStruct

func (h heapInit) Len() int           { return len(h) }
func (h heapInit) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h heapInit) Swap(i, j int)      {
	if len(h) < 2 {
		return
	}
	h[i], h[j] = h[j], h[i]
}

func (h *heapInit) Push(x any) {
		*h = append(*h, x.(heapStruct))
	}

func (h *heapInit) Pop() any {
	if len(*h) == 0 {
		return heapStruct{-1}
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var (
	maxInt int = 2147483647
	matrix [][]int
)

func Btoi(b bool) int {
	if b == true {return 1}
	return 0
}

func indexOfMin(from int, last int, vis []bool, remain int) int {
	min, idx := maxInt, 0

	if remain == 1 {
		idx = last * Btoi(matrix[from][last] != 0)
	}
	for i := 1; i < len(vis); i++ {
		if i == from || i == last {
			continue
		}
		if !vis[i] && matrix[from][i] < min && matrix[from][i] != 0{
			min = matrix[from][i]
			idx = i
		}
	}
	if idx == 0 {
		return -1
	}
	return idx
}

func getAngryLowWay(apex int) int {
	minDist := -1
	remain := len(matrix[1]) - 1
	visited := make([]bool, len(matrix[1]))
	idx := apex
	wayLen := 0
	for true {
		newIdx := indexOfMin(idx, apex, visited, remain)
		if newIdx == -1 {
			break
		}
		remain--
		visited[newIdx] = true
		wayLen += matrix[idx][newIdx]
		if remain == 0 {
			minDist = wayLen
		}
		idx = newIdx
	}
	return minDist
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	apexCount, _ := strconv.Atoi(scn.Text())
	ans := 0
	matrix = make([][]int, apexCount + 1)
	sumBest := 0
	for i := 1; i <= apexCount; i++ {
		matrix[i] = make([]int, apexCount + 1)
		minInLine := maxInt
		for j := 1; j <= apexCount; j++ {
			scn.Scan()
			value, _ := strconv.Atoi(scn.Text())
			if value != 0 && minInLine > value {
				minInLine = value
			}
			matrix[i][j] = value
		}
		matrix[i][0] = minInLine
		sumBest += minInLine
	}
	inputFile.Close()
	if apexCount != 1 {
		ans = getAngryLowWay(1)
		ans = getMinimalWay(1, ans, sumBest)
	}
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strconv.Itoa(ans)))
	outputFile.Close()
}

func getMinimalWay(apex int, bestWay int, sumBest int) int {
	if bestWay == -1 {
		return -1
	}
	bufferArr := make([]int, 3 + len(matrix[1]) + 1, 3 + len(matrix[1]) + 1)
	bufferArr[0] = 0
	bufferArr[1] = sumBest
	bufferArr[2] = apex
	bufferArr[len(bufferArr)-2] = len(matrix[1]) - 1
	h := &heapInit{bufferArr}
	heap.Init(h)
	minDist := bestWay
	for true {
		fromHeap := heap.Pop(h).(heapStruct)
		if fromHeap[len(fromHeap)-1] == 1 {
			if minDist > fromHeap[0] {
				minDist = fromHeap[0]
			}
			break
		}
		if fromHeap[0] == -1 {
			break
		}
		if fromHeap[0] + fromHeap[1] >= bestWay {
			continue
		}
		idx := fromHeap[2]
		for i, v := range matrix[idx] {
			if i != 0 && i!= apex && v != 0 && !fromHeap.CheckVis(i) {
				newHS := heapStruct(append(make([]int, 0, 3 + len(matrix[1]) + 1), fromHeap...))
				newHS.Add(i,v)
				heap.Push(h, newHS)
			}
			if fromHeap[len(fromHeap)-2] == 1 {
				newHS := heapStruct(append(make([]int, 0, 3 + len(matrix[1]) + 1), fromHeap...))
				newHS.Add(apex,matrix[idx][apex])
				heap.Push(h, newHS)
				break
			}
		}
	}
	return minDist
}