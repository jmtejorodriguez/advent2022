package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "10/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	instructions := bytes.Split(dat, []byte("\n"))

	strength := 0
	currentCycles := 0
	register := 1
	nextMeasurement := 20
	for _, instruction := range instructions {
		if len(instruction) == 0 {
			continue
		}
		var cycles int
		delta := 0
		if bytes.HasPrefix(instruction, []byte("noop")) {
			cycles = 1
		} else {
			cycles = 2
			delta, _ = strconv.Atoi(string(instruction[5:]))
		}
		currentCycles += cycles
		if currentCycles >= nextMeasurement {
			strength += nextMeasurement * register
			nextMeasurement += 40
		}
		register += delta
	}

	fmt.Println(strength)
}
