package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type heapInit [][]float64

func (h heapInit) Len() int           { return len(h) }
func (h heapInit) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h heapInit) Swap(i, j int)      {
	if len(h) < 2 {
		return
	}
	h[i], h[j] = h[j], h[i] }

func (h *heapInit) Push(x any) {
	*h = append(*h, x.([]float64))
}

func (h *heapInit) Pop() any {
	if len(*h) == 0 {
		return []float64{-1}
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var (
	maxF64 float64 = math.MaxFloat64
	maxInt int = math.MaxInt32
	cityCount int
	cityParam [][]int
	list [][][]int
	distance [][]float64
	timeToCap []float64
)


func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	cityCount, _ = strconv.Atoi(scn.Text())
	cityParam = make([][]int, cityCount + 1)
	for i := 1 ; i <= cityCount; i++ {
		scn.Scan()
		time, _ := strconv.Atoi(scn.Text())
		scn.Scan()
		speed, _ := strconv.Atoi(scn.Text())
		cityParam[i] = []int{time, speed}
	}
	list = make([][][]int, cityCount + 1)
	for i := 1; i < cityCount; i++ {
		scn.Scan()
		from, _ := strconv.Atoi(scn.Text())
		scn.Scan()
		to, _ := strconv.Atoi(scn.Text())
		scn.Scan()
		dist, _ := strconv.Atoi(scn.Text())
		list[from] = append(list[from], []int{to, dist})
		list[to] = append(list[to], []int{from, dist})
	}
	inputFile.Close()
	distance = make([][]float64, cityCount + 1)
	for i := 1; i <= cityCount; i++ {
		distance[i] = make([]float64, cityCount + 1)
		for j := 1; j <= cityCount; j++ {
			distance[i][j] = maxF64
		}
	}
	for i := 1; i <= cityCount; i++ {
		distance[i][i] = 0
		h := &heapInit{[]float64{0, float64(i)}}
		heap.Init(h)
		for true {
			fromHeap := heap.Pop(h).([]float64)
			dist := fromHeap[0]
			if dist == -1 {
				break
			}
			idx := int(fromHeap[1])
			for _, v := range list[idx] {
				if distance[i][v[0]] > distance[i][idx] + float64(v[1]) {
					distance[i][v[0]] = distance[i][idx] + float64(v[1])
					heap.Push(h, []float64{distance[i][v[0]], float64(v[0])})
				}
			}
		}
	}
	timeToCap = make([]float64, cityCount + 1)
	for i := 1; i <= cityCount; i++ {
		timeToCap[i] = maxF64
	}
	timeToCap[1] = 0
	for i, v := range list {
		if i > 1 && timeToCap[i] > float64(cityParam[i][0]) + distance[i][1]/float64(cityParam[i][1]) {
			timeToCap[i] = float64(cityParam[i][0]) + distance[i][1]/float64(cityParam[i][1])
		}
		for _, innV := range v {
			currentRebase := float64(cityParam[innV[0]][0])
			currentSpeed := float64(cityParam[innV[0]][1])
			currentDistToCap := distance[innV[0]][1]
			if timeToCap[innV[0]] > currentRebase + currentDistToCap/currentSpeed {
				timeToCap[innV[0]] = currentRebase + currentDistToCap/currentSpeed
			} 
			update(i, i)
		}
	}
	var minTime float64 = 0
	minI := 0
	for i := 2; i <= cityCount; i++ {
		if timeToCap[i] > minTime {
			minTime = timeToCap[i] 
			minI = i
		}
	}
	var way []int
	minTime, way = getMinTime(minI, 1)
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(fmt.Sprintf("%.10f\n%s", minTime, strings.Trim(fmt.Sprintf("%v", way), "[]"))))
	outputFile.Close()
}

func update(where, who int) {
	for _, v := range list[who] {
		if timeToCap[v[0]] > float64(cityParam[v[0]][0]) + distance[v[0]][where]/float64(cityParam[v[0]][1]) +  timeToCap[where]{
			timeToCap[v[0]] = float64(cityParam[v[0]][0]) + distance[v[0]][where]/float64(cityParam[v[0]][1]) +  timeToCap[where]
			update(where, v[0])
		} else if timeToCap[where] > float64(cityParam[where][0]) + distance[where][v[0]]/float64(cityParam[where][1]) + timeToCap[v[0]]{ // пересадка в ребенке
			timeToCap[where] = float64(cityParam[where][0]) + distance[where][v[0]]/float64(cityParam[where][1]) + timeToCap[v[0]]
			update(v[0], where)
		}
	}
}

func getMinTime(from, to int) (float64, []int) {
	speedArr := make([]float64, cityCount + 1)
	h := &heapInit{[]float64{0, 0, float64(from), float64(from)}}
	heap.Init(h)
	minTime := maxF64
	var way []float64
	for true {
		fromHeap := (heap.Pop(h).([]float64))
		time := fromHeap[0]
		if time == -1 {
			break
		}
		idx := int(fromHeap[2])
		if idx == to && time < minTime {
			minTime = time
			way = fromHeap[3:]
			continue
		}
		if time > minTime {
			continue
		}
		currentSpeed := fromHeap[1]
		if speedArr[idx] > currentSpeed {
			continue
		}
		speedArr[idx] = currentSpeed
		for _, v := range list[idx] {
			if speedArr[v[0]] != 0 && speedArr[v[0]] >= currentSpeed && speedArr[v[0]] >= float64(cityParam[idx][1]){
				continue
			}
			innerSpeed := currentSpeed
			newTime := maxF64
			withTransfer := maxF64
			innerWay := fromHeap[3:]
			if innerSpeed < float64(cityParam[idx][1]) {
				withTransfer = time + float64(cityParam[idx][0]) + float64(v[1])/float64(cityParam[idx][1])
				newSpeed := float64(cityParam[idx][1])
				w := append(innerWay, float64(idx))
				out := append([]float64{withTransfer, newSpeed, float64(v[0])}, w...)
				heap.Push(h, out)
			}
			if innerSpeed != 0{
				newTime = time + float64(v[1])/float64(innerSpeed)
			}
			if withTransfer < newTime {
				continue
			}
			out := append([]float64{newTime, innerSpeed, float64(v[0])}, innerWay...)
			heap.Push(h, out)
		}
	}
	ans := make([]int, len(way))
	for i := 0; i < len(way) - 1; i ++ {
		ans[i] = int(way[i + 1])
	}
	ans[len(ans) - 1] = 1
	return minTime, ans
}