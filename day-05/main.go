package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type inst = int
type stacks map[inst][]string
type instruction [3]inst

const (
	instMove inst = iota
	instFrom
	instTo
)

func (s *stacks) String() string {
	return fmt.Sprintf("%+v", (*s))
}

func lexCrates(input string) stacks {

	var output = make(map[inst][]string)
	curCol := 1
	var blankCount int

	for _, ch := range input {
		if unicode.IsLetter(ch) {
			output[inst(curCol)] = append(output[inst(curCol)], string(ch))
			blankCount = 0
			curCol++
		}
		if ch == ' ' {
			blankCount++
			if blankCount >= 4 {
				curCol++
				blankCount = 0
			}
		}
		if ch == '\n' {
			curCol = 1
			blankCount = 0
		}
		if unicode.IsNumber(ch) {
			break
		}
	}

	for _, v := range output {
		slices.Reverse(v)
	}

	return output
}

func (s *stacks) readInstruction(line string) instruction {

	instArr := instruction{}
	cur := 0
	var strNum string
	for _, ch := range line {
		if unicode.IsNumber(ch) {
			strNum += string(ch)
		}
		if ch == ' ' && strNum != "" {
			num, _ := strconv.Atoi(strNum)
			instArr[cur] = num
			cur++
			strNum = ""
		}
	}

	num, _ := strconv.Atoi(strNum)
	instArr[cur] = num

	return instArr
}

func (s *stacks) writeInstruction(inst instruction) {

	for range int(inst[instMove]) {
		from := (*s)[inst[instFrom]]
		to := (*s)[inst[instTo]]
		val := from[len(from)-1]
		to = append(to, val)
		(*s)[inst[instFrom]] = from[:len(from)-1]
		(*s)[inst[instTo]] = to
	}

}

func (s *stacks) writeBatchInstruction(inst instruction) {

	from := (*s)[inst[instFrom]]
	to := (*s)[inst[instTo]]
	val := from[len(from)-inst[instMove]:]
	to = append(to, val...)
	(*s)[inst[instFrom]] = from[:len(from)-inst[instMove]]
	(*s)[inst[instTo]] = to

}

// return the start index of the instructions
func findInstructions(input string) int {

	var idx int

	for i, ch := range input {
		if unicode.IsLetter(ch) && unicode.IsLower(ch) {
			idx = i
			break
		}
	}

	return idx

}

func (s *stacks) peekAll() string {
	var res string
	key := 1

	for len((*s)[key]) > 0 {
		res += (*s)[key][len((*s)[key])-1]
		key++
	}

	return res
}

func main() {

	// use input or testInput for a smaller example
	curInput := input

	s := lexCrates(curInput)
	idx := findInstructions(curInput)
	lines := strings.SplitSeq(curInput[idx:], "\n")

	fmt.Println(s)

	for line := range lines {
		inst := s.readInstruction(line)
		// part 1:
		// s.writeInstruction(inst)
		// part 2:
		s.writeBatchInstruction(inst)
	}
	fmt.Println(s.peekAll())

}
