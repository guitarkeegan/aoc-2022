package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var dbg = func() func(format string, args ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(format string, args ...any) {}
	}
	return func(format string, args ...any) {
		fmt.Printf(format+"\n", args...)
	}
}()

type comm struct {
	input  string
	cur    map[byte]bool
	lPos   int
	rPos   int
	target int
}

func (c *comm) String() string {
	in := c.input[:c.rPos]
	js, err := json.MarshalIndent(c.cur, "", "  ")
	if err != nil {
		log.Fatalf("json err: %v\n", err)
	}
	return fmt.Sprintf("INPUT: %s\nCUR: %s\nL_POS: %d\nR_POS: %d\nTARGET: %d\n", in, string(js), c.lPos, c.rPos, c.target)
}

func NewComm(input string, target int) *comm {

	if len(input) < 4 {
		return nil
	}
	co := make(map[byte]bool)
	co[input[0]] = true
	return &comm{
		input:  input,
		cur:    co,
		rPos:   1,
		target: target,
	}
}

// find the index of the last valid character in the sequence
// return -1 if no valid start is found
func (c *comm) FindPacketStart() int {

	dbg("FindStart: %+v", c.cur)

	for c.rPos < len(c.input) {
		// check if char is in set
		newChar := c.input[c.rPos]
		dbg("  newChar: %d, map: %+v", newChar, c.cur)
		if _, ok := c.cur[newChar]; ok {
			dbg("  have char already...")
			// if it is, remove from left until it is not, then add it
			for c.cur[newChar] && c.lPos <= c.rPos {
				dbg("  delete")
				delete(c.cur, c.input[c.lPos])
				dbg("  map: %+v", c.cur)
				c.lPos++
			}
		}
		c.cur[newChar] = true
		c.rPos++
		// if it isn't, add it to the set, and go forward
		// if the target == len(cur), return c.rPos
		if c.target == c.rPos-c.lPos {
			// 1 indexed kara, return the previous pos
			dbg("%s", c)
			return c.rPos
		}
	}

	return -1
}

func (c *comm) FindMsgStart() int {

	dbg("FindMsgStart: %+v", c.cur)

	for c.rPos < len(c.input) {
		newChar := c.input[c.rPos]
		if _, ok := c.cur[newChar]; ok {
			for c.cur[newChar] && c.lPos <= c.rPos {
				delete(c.cur, c.input[c.lPos])
				c.lPos++
			}
		}
		c.cur[newChar] = true
		c.rPos++
		if c.target == c.rPos-c.lPos {
			return c.rPos
		}
	}

	return -1
}

func main() {

	comm := NewComm(input, 4)
	comm2 := NewComm(input, 14)
	res := comm.FindPacketStart()
	res2 := comm2.FindMsgStart()

	fmt.Println(res)
	fmt.Println(res2)
}
