package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	capacity, _ := strconv.Atoi(scn.Text())
	scn.Scan()
	size, _ := strconv.Atoi(scn.Text())
	floors := make([]int, size)
	lastFiled := 0
	for i:=0; i<size;i++ {
		scn.Scan()
		value, _ := strconv.Atoi(scn.Text())
		if value != 0 {
			lastFiled = i
		}
		floors[i] = value
	}
	inputFile.Close()
	wastedTime := 0
	currentFill := 0
	for true {
		if currentFill != 0 {
			currentFill += floors[lastFiled]
			floors[lastFiled] = 0
			if currentFill >= capacity {
				floors[lastFiled] = (currentFill/capacity - 1) * capacity + currentFill % capacity
				currentFill = capacity
			}
		} else {
			if floors[lastFiled]/capacity != 0 {
				wastedTime += floors[lastFiled]/capacity * (lastFiled + 1) * 2
			}
			if floors[lastFiled]%capacity != 0 {
				wastedTime += lastFiled + 1
			}
			currentFill = floors[lastFiled]%capacity
			floors[lastFiled] = 0
		}
		trigger := false
		for i := lastFiled ; i >= 0; i-- {
			if floors[i] != 0 {
				trigger = true
				wastedTime += lastFiled - i
				lastFiled = i
				break
			}
		}
		if currentFill == capacity {
			wastedTime += lastFiled + 1
			currentFill = 0
			continue
		}
		if !trigger {
			if currentFill != 0 {
				wastedTime += lastFiled + 1
			}
			break
		}
		
	}

	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(fmt.Sprintf("%d\n", wastedTime)))
	outputFile.Close()
}
