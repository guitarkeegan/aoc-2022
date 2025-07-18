package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var dbg = func() func(format string, as ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(format string, as ...any) {}
	}
	return func(format string, as ...any) {
		fmt.Printf(format+"\n", as...)
	}
}()

type Clock struct {
	cycle   int
	x       int
	reg     int
	signals []int
}

func (c *Clock) String() string {
	return fmt.Sprintf("cycle: %d, x: %d, reg: %d, signals: %+v", c.cycle, c.x, c.reg, c.signals)
}

func NewClock() *Clock {
	return &Clock{
		cycle: 0,
		x:     1,
	}
}

func Signal(cycle, x int) (int, bool) {

	signals := [6]int{20, 60, 100, 140, 180, 220}

	for _, n := range signals {
		if cycle == n {
			return cycle * n, true
		}
	}

	return 0, false
}

func (c *Clock) runInst(inst string) {
	line := strings.Split(strings.TrimSpace(inst), " ")

	if c.reg > 0 {

		c.x += c.reg
		c.reg = 0
		c.cycle++
		if sig, ok := Signal(c.cycle, c.x); ok {
			c.signals = append(c.signals, sig)
		}
		dbg("  proc: %v", c)
	}

	// noop
	if len(line) == 1 {
		// op := line[0]
		c.x += c.reg
		c.reg = 0
		c.cycle++
		if sig, ok := Signal(c.cycle, c.x); ok {
			c.signals = append(c.signals, sig)
		}
		dbg("  noop: %v", c)
		return
	}

	// addx
	nxtReg, err := strconv.Atoi(line[1])
	if err != nil {
		log.Fatalf("not a num %v\n", err)
	}

	c.x += c.reg
	c.reg = nxtReg
	c.cycle++
	if sig, ok := Signal(c.cycle, c.x); ok {
		c.signals = append(c.signals, sig)
	}
	dbg("  addx: %v", c)
}

func main() {

	dbg("start program")

	lines := strings.SplitSeq(lgTestInput, "\n")
	c := NewClock()

	dbg("clock: %v", c)
	for line := range lines {
		c.runInst(line)
	}
	dbg("final inst")
	c.x += c.reg
	c.reg = 0
	c.cycle++

	fmt.Printf("final: %v\n", c)
}
