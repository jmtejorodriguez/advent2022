package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "1/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}
	text := string(dat)

	lines := strings.Split(text, "\n")

	var qty = 0
	var max = 0
	for _, s := range lines {
		if s == "" {
			max = Max(max, qty)
			qty = 0
		} else {
			add, _ := strconv.Atoi(s)
			qty += add
		}
	}
	max = Max(max, qty)
	fmt.Println(max)
}
