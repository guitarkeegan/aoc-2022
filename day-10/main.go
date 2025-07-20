package main

import (
	"fmt"
	"iter"
	"log"
	"os"
	"slices"
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

var dbg2 = func() func(format string, as ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(format string, as ...any) {}
	}
	return func(format string, as ...any) {
		fmt.Printf(format+"\n", as...)
	}
}()

type Clock struct {
	cycle       int
	x           int
	reg         int
	signals     map[int]struct{}
	signalsSize int
	next        func() (line string, ok bool)
	stop        func()
	screen      string
	pos         int
}

type op int

const (
	noop op = iota
	addx
)

func (c *Clock) Draw() {

	c.pos = c.cycle % 40
	// newline?
	if c.pos == 39 {
		c.screen += "\n"
	}
	// if one of the 3 sprit pixes is within the
	// cur, draw the pixel #, else draw .
	if c.cycle == c.x ||
		c.cycle == c.x-1 ||
		c.cycle == c.x+1 {
		c.screen += "#"
	} else {
		c.screen += "."
	}
}

func (c *Clock) String() string {
	return fmt.Sprintf("cycle: %d, x: %d, reg: %d, signals: %+v", c.cycle, c.x, c.reg, c.signals)
}

func NewClock(program iter.Seq[string]) *Clock {
	next, stop := iter.Pull(program)

	return &Clock{
		cycle:   0,
		x:       1,
		next:    next,
		stop:    stop,
		signals: make(map[int]struct{}),
	}
}

func (c *Clock) Tick() {
	c.cycle++
}

func (c *Clock) Signal(cycle, x int) (int, bool) {

	signals := []int{20, 60, 100, 140, 180, 220}

	if slices.Contains(signals, cycle) {
		return cycle * x, true
	}

	return 0, false
}

func (c *Clock) addx() {
	c.x += c.reg
}

func is(line string) (op, int) {

	parts := strings.Split(strings.TrimSpace(line), " ")
	if len(parts) == 1 {
		return noop, 0
	}
	v, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("not int: %q\n", err)
	}
	return addx, v

}

// if there is a problem, return an error
// return false until all instructions have be executed
func (c *Clock) Run() (error, bool) {

	// before the next cycle starts
	// addx
	if c.reg != 0 {
		c.Tick()
		dbg("RunStart: %s", c)
		if sig, ok := c.Signal(c.cycle, c.x); ok {
			c.signals[sig] = struct{}{}
		}
		c.Draw()
		c.addx()
		c.reg = 0
		dbg("  RunEnd addx: %s", c)
		return nil, false
	}

	c.Tick()
	dbg("RunStart: %s", c)

	if sig, ok := c.Signal(c.cycle, c.x); ok {
		c.signals[sig] = struct{}{}
	}
	c.Draw()

	line, ok := c.next()
	if !ok {
		c.stop()
		dbg("  RunEnd Complete: %s", c)
		return nil, true
	}

	op, value := is(line)

	switch op {
	case noop:
	case addx:
		c.reg = value
	}

	dbg("  RunEnd: %s", c)
	return nil, false

}

func main() {

	dbg("start program")

	lines := strings.SplitSeq(lgTestInput, "\n")
	c := NewClock(lines)

	for {
		err, ok := c.Run()
		if err != nil {
			log.Fatalf("error: %q\n", err)
		}
		if ok {
			break
		}
	}

	fmt.Printf("final: %v\n", c)
	var total int
	for x := range c.signals {
		total += x
	}
	fmt.Printf("Part 1: total: %d\n", total)
	fmt.Printf("Part 2: \n%s", c.screen)
}
