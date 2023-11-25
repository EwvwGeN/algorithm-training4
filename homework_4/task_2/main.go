package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	tableSize, _ := strconv.Atoi(scn.Text())
	onLine := make([]bool, tableSize*2)
	onDiagonal := make([]bool, 2*(2*tableSize - 1))
	ans := setFigure(onLine, onDiagonal, tableSize, 0, 0, 0, false)
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strconv.Itoa(ans)))
	outputFile.Close()
}

func Btoi(b bool) int {
	if b {return 1}
	return 0
}

func setFigure(lines, diagonals []bool, size int, currentCount int, x, y int, set bool) int {
	var res int
	countF := currentCount
	if countF >= size {
		return 1
	}
	if set {
		lines[x] = true
		lines[size + y] = true
		diagonals[x+y] = true
		diagonals[2*size-1+int(math.Abs(float64(x - y - size + 1)))] = true
	}
	for i := x; i < size; i++ {
		if lines[i]{
			continue
		}
		if size - i < size - countF {
			return res
		}
		for j := y*Btoi(i==x); j < size; j++ {
			if lines[size + j]{
				continue
			}
			
			if !diagonals[i+j] && !diagonals[2*size-1+int(math.Abs(float64(i - j - size + 1)))] {
				newL := append(make([]bool,0, len(lines)), lines...)
				newD := append(make([]bool,0, len(diagonals)), diagonals...)
				res += setFigure(newL, newD, size, countF + 1, i, j, true)
			}
		}
	}
	return res
}