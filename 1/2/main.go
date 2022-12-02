package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

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
	var qties []int
	for _, s := range lines {
		if s == "" {
			qties = append(qties, qty)
			qty = 0
		} else {
			add, _ := strconv.Atoi(s)
			qty += add
		}
	}
	qties = append(qties, qty)
	sort.Ints(qties)
	var result = 0
	for i := len(qties) - 1; i >= 0 && i > len(qties)-4; i-- {
		result += qties[i]
	}
	fmt.Println(result)
}
