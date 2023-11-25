package main

import (
	"bufio"
	"os"
)

type Stack struct {
	s []rune
	size int
	filed int
}

func (s *Stack) IsEmpty() bool {
	if s.filed == s.size {
		return false
	}
	return true
}

func (s *Stack) Push(char rune) int {
	if s.filed == s.size {
		return 1
	}
	s.s[s.filed] = char
	s.filed++
	return 0
}

func (s *Stack) Pop() (rune, int) {
	if s.filed == 0 {
		return 0, 1
	}
	s.filed--
	tmp := s.s[s.filed]
	s.s[s.filed] = 0
	return tmp, 0
}

func main() {
	inputFile, _ := os.Open("input.txt")
	scn := bufio.NewScanner(inputFile)
	buf := make([]byte, 1024*1024)
	scn.Buffer(buf, 1024*1024)
	scn.Scan()
	str := scn.Text()
	inputFile.Close()
	b := checkStr(str)
	ans := "yes"
	if b == false {
		ans = "no"
	}
	outputFile, _ := os.Create("output.txt")
	outputFile.Write([]byte(ans))
	outputFile.Close()
}

func checkStr(str string) bool {
	if len(str) % 2 != 0 {
		return false
	}
	
	stack := Stack{
		s: make([]rune, len(str)/2, len(str)/2),
		size: len(str)/2,
	}

	for _, v := range str {
		switch v {
			case '(', '[', '{' :{
				if stack.Push(v) == 1 {
					return false
				}
			}
			case ')': {
				if r, err := stack.Pop(); err == 1 || r != 40 {
					return false
				}
			}
			case ']': {
				if r, err := stack.Pop(); err == 1 || r != 91 {
					return false
				}
			}
			case '}': {
				if r, err := stack.Pop(); err == 1 || r != 123 {
					return false
				}
			}
		}
	}

	return true
}