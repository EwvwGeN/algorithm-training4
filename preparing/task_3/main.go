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
	coordinates := make([]float64, 4)
	for i := 0; i < 4; i++ {
		scn.Scan()
		value, _ := strconv.ParseFloat(scn.Text(), 64)
		coordinates[i] = value
	}
	inputFile.Close()

	firstLen := math.Sqrt(coordinates[0]*coordinates[0] + coordinates[1]*coordinates[1])
	secondLen := math.Sqrt(coordinates[2]*coordinates[2] + coordinates[3]*coordinates[3])
	len := firstLen + secondLen
	r := min(firstLen, secondLen)
	if firstLen > secondLen {
		coordinates[0], coordinates[2] = coordinates[2], coordinates[0]
		coordinates[1], coordinates[3] = coordinates[3], coordinates[1]
	}

	angle := math.Abs(math.Atan2(
		coordinates[3],
		coordinates[2],
		) - 
		math.Atan2(
			coordinates[1],
			coordinates[0],
		))

	if angle == 0 && len != r*2 {
		len = math.Abs(firstLen - secondLen)
	}

	if angle < 2 && angle != 0{
		yPoint := coordinates[3]
		xPoint := coordinates[2]
		if yPoint != 0 {
			yPoint = r*coordinates[3] / math.Sqrt(coordinates[3]*coordinates[3] + coordinates[2] * coordinates[2])
			xPoint = yPoint*coordinates[2] / coordinates[3]
		}
		
		
		arcLen := angle*r
		remLen := math.Sqrt(math.Pow(coordinates[3] - yPoint, 2) + math.Pow(coordinates[2] - xPoint, 2))

		len = arcLen + remLen
	}

	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(strconv.FormatFloat(len, 'f', 12, 64)))
	outputFile.Close()
}
