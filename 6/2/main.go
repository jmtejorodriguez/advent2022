package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "6/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	last := 0
	for i := 1; i < len(dat); i++ {
		for j := i - 1; j >= last; j-- {
			if dat[j] == dat[i] {
				last = j + 1
				break
			}
		}
		if i-last == 13 {
			fmt.Print(i + 1)
			break
		}
	}
}
