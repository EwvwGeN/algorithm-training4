package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type heapInit [][]int64

func (h heapInit) Len() int           { return len(h) }
func (h heapInit) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h heapInit) Swap(i, j int)      {
	if len(h) < 2 {
		return
	}
	h[i], h[j] = h[j], h[i] }

func (h *heapInit) Push(x any) {
	*h = append(*h, x.([]int64))
}

func (h *heapInit) Pop() any {
	if len(*h) == 0 {
		return int64(-1)
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x[1]
}

var (
	maxUInt int64 = 9223372036854775807
)


func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	apexCount, _ := strconv.Atoi(scn.Text())
	list := make([][][]int, apexCount + 1)
	scn.Scan()
	roadCount, _ := strconv.Atoi(scn.Text())
	for i := 0; i < roadCount; i++ {
		scn.Scan()
		from, _ := strconv.Atoi(scn.Text())
		scn.Scan()
		to, _ := strconv.Atoi(scn.Text())
		scn.Scan()
		dist, _ := strconv.Atoi(scn.Text())
		list[from] = append(list[from], []int{to, dist})
		list[to] = append(list[to], []int{from, dist})
	}
	scn.Scan()
	from, _ := strconv.Atoi(scn.Text())
	scn.Scan()
	to, _ := strconv.Atoi(scn.Text())
	inputFile.Close()
	distance := make([]int64, apexCount + 1)
	for i := 0; i < len(distance); i++ {
		distance[i] = maxUInt
	}
	distance[from] = 0
	h := &heapInit{[]int64{0, int64(from)}}
	heap.Init(h)
	minDist := int64(-1)
	for true {
		idx := heap.Pop(h).(int64)
		if idx == -1 {
			break
		}
		for _, v := range list[idx] {
			if distance[v[0]] > distance[idx] + int64(v[1]) {
				distance[v[0]] = distance[idx] + int64(v[1])
				heap.Push(h, []int64{distance[v[0]], int64(v[0])})
			}
		}
		if idx == int64(to) {
			minDist = distance[idx]
			break
		}
	}

	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(fmt.Sprintf("%d", minDist)))
	outputFile.Close()
}