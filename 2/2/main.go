package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var values = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "2/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}
	text := string(dat)

	lines := strings.Split(text, "\n")

	var sum = 0
	for _, s := range lines {
		if s == "" {
			continue
		}
		sum += values[s]
	}
	fmt.Println(sum)
}
