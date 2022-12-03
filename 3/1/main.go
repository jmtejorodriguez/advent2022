package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func toPriority(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b - 'a' + 1)
	}
	return int(b - 'A' + 27)
}

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "3/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	lines := bytes.Split(dat, []byte("\n"))

	var sum = 0
	for _, s := range lines {
		var found = map[byte]bool{}
		for i := 0; i < len(s)/2; i++ {
			found[s[i]] = true
		}
		for i := len(s) / 2; i < len(s); i++ {
			if found[s[i]] {
				sum += toPriority(s[i])
				break
			}
		}
	}
	fmt.Println(sum)
}
