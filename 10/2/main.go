package main

import (
	"bytes"
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
	currentInstruction := 0
	var completionTime int
	if bytes.HasPrefix(instructions[0], []byte("noop")) {
		completionTime = 1
	} else {
		completionTime = 2
	}
	cycle := 0
	register := 1
	for {
		position := cycle % 40
		if position == 0 {
			println()
		}
		if position >= register-1 && position <= register+1 {
			print("#")
		} else {
			print(".")
		}
		cycle++
		completionTime--
		if completionTime == 0 {
			if bytes.HasPrefix(instructions[currentInstruction], []byte("addx")) {
				delta, _ := strconv.Atoi(string(instructions[currentInstruction][5:]))
				register += delta
			}
			currentInstruction++
			if bytes.HasPrefix(instructions[currentInstruction], []byte("noop")) {
				completionTime = 1
			} else {
				completionTime = 2
			}
		}
		if cycle == 240 {
			break
		}
	}
}
