package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	size int
	remain int
	filed int
}

func (s *Stack) IsEmpty() bool {
	if s.remain == s.size {
		return true
	}
	return false
}

func (s *Stack) Push() int {
	if s.remain == 0 {
		return 1
	}
	s.filed++
	s.remain--
	return 0
}

func (s *Stack) Pop() int {
	if s.remain == s.size {
		return 1
	}
	s.filed--
	s.remain++
	return 0
}

func main() {
	inputFile, _ := os.Open("brackets2.in")
	scn := bufio.NewScanner(inputFile)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	inputLen, _ := strconv.Atoi(scn.Text())
	inputFile.Close()
	var ans []string
	if inputLen % 2 == 0 {
		fistStack, secondStack := &Stack{
			remain: inputLen/2,
			size: inputLen/2,
		}, &Stack{
			remain: inputLen/2,
			size: inputLen/2,
		}
		ans = getCorrectBracketSequencen(fistStack, secondStack, "", "", inputLen)
	}
	outputFile, _ := os.Create("brackets2.out")
	outputFile.Write([]byte(strings.Join(ans, "\n")))
	outputFile.Close()
}

func getCorrectBracketSequencen(stackFirst, stackSecond *Stack, str string, open string, maxLen int) []string {
	var out []string
	for _, v := range "([)]" {
		switch v {
			case '(':  {
				str = str + string(v)
				if stackSecond.remain == 0 || stackFirst.Push() == 1{
					str = str[:len(str)-1]
					continue
				}
				if len(str) == maxLen || len(str) + stackFirst.filed + stackSecond.filed > maxLen  {
					stackFirst.Pop()
					str = str[:len(str)-1]
					continue
				}
				open = open + string(v)
				buf := getCorrectBracketSequencen(stackFirst, stackSecond, str, open, maxLen)
				out = append(out, buf...)
				open = open[:len(open)-1]
				str = str[:len(str)-1]
				stackFirst.Pop()
			}
			case '[': {
				str = str + string(v)
				if stackFirst.remain == 0 || stackSecond.Push() == 1 {
					str = str[:len(str)-1]
					continue
				}
				if len(str) == maxLen || len(str) + stackFirst.filed + stackSecond.filed > maxLen {
					stackSecond.Pop()
					str = str[:len(str)-1]
					continue
				}
				open = open + string(v)
				buf := getCorrectBracketSequencen(stackFirst, stackSecond, str, open, maxLen)
				out = append(out, buf...)
				open = open[:len(open)-1]
				str = str[:len(str)-1]
				stackSecond.Pop()
			}
			case ')': {
				if len(open) < 1 || open[len(open)-1] == '[' {
					continue
				}
				str = str + string(v)
				if stackSecond.remain == 0 || stackFirst.Pop() == 1  {
					str = str[:len(str)-1]
					continue
				}
				if str[len(str) - 2] == '[' || len(str) + stackFirst.filed + stackSecond.filed > maxLen {
					str = str[:len(str)-1]
					stackFirst.Push()
					continue
				}
				if len(str) == maxLen{
					if stackFirst.IsEmpty() && stackSecond.IsEmpty()  {
						out = append(out, str)
						stackFirst.Push()
						return out
					}
				}
				tmp := open[len(open)-1]
				open = open[:len(open)-1]
				buf := getCorrectBracketSequencen(stackFirst, stackSecond, str, open, maxLen)
				out = append(out, buf...)
				open = open + string(tmp)
				str = str[:len(str)-1]
				stackFirst.Push()
			}
			case ']': {
				if len(open) < 1 || open[len(open)-1] == '(' {
					continue
				}
				str = str + string(v)
				if stackFirst.remain == 0 || stackSecond.Pop() == 1 {
					str = str[:len(str)-1]
					continue
				}
				if str[len(str) - 2] == '(' || len(str) + stackFirst.filed + stackSecond.filed > maxLen {
					str = str[:len(str)-1]
					stackSecond.Push()
					continue
				}
				if len(str) == maxLen {
					if stackFirst.IsEmpty() && stackSecond.IsEmpty()  {
						out = append(out, str)
						stackSecond.Push()
						return out
					}
				}
				tmp := open[len(open)-1]
				open = open[:len(open)-1]
				buf := getCorrectBracketSequencen(stackFirst, stackSecond, str, open, maxLen)
				out = append(out, buf...)
				open = open + string(tmp)
				str = str[:len(str)-1]
				stackSecond.Push()
			}
		}
	}
	return out
}
