package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "5/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	lines := bytes.Split(dat, []byte("\n"))

	var bottom = 0

	for ; !bytes.HasPrefix(lines[bottom], []byte(" 1 ")); bottom++ {

	}
	// I assume stacks are numbered from 1 to N and ordered in the input from left to right
	// I assume at least one stack
	var stackQty = 1
	for i := 0; i*4 < len(lines[bottom]); i++ {
		stackQty++
	}

	var stacks [][]byte
	for i := 0; i < stackQty; i++ {
		stacks = append(stacks, []byte{})
	}
	for i := bottom - 1; i >= 0; i-- {
		row := lines[i]
		for j := 0; j*4+1 < len(row); j++ {
			ch := row[j*4+1]
			if ch == ' ' {
				continue
			}
			stacks[j] = append(stacks[j], ch)
		}
	}
	moveRegexp := regexp.MustCompile(`move (?P<Quantity>\d+) from (?P<Origin>\d+) to (?P<Destination>\d+)`)
	for i := bottom + 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		matches := moveRegexp.FindSubmatch(lines[i])
		qty, _ := strconv.Atoi(string(matches[moveRegexp.SubexpIndex("Quantity")]))
		origin, _ := strconv.Atoi(string(matches[moveRegexp.SubexpIndex("Origin")]))
		origin--
		destination, _ := strconv.Atoi(string(matches[moveRegexp.SubexpIndex("Destination")]))
		destination--
		for j := 0; j < qty; j++ {
			index := len(stacks[origin]) - 1
			stacks[destination] = append(stacks[destination], stacks[origin][index])
			stacks[origin] = stacks[origin][:index]
		}
	}

	for i := 0; i < len(stacks); i++ {
		if len(stacks[i]) == 0 {
			continue
		}
		fmt.Print(string(stacks[i][len(stacks[i])-1]))
	}
	fmt.Println()
}
