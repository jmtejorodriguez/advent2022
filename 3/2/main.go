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
	for i := 0; i+2 < len(lines); i += 3 {
		s1 := lines[i]
		s2 := lines[i+1]
		s3 := lines[i+2]
		var foundOnce = map[byte]bool{}
		for i := 0; i < len(s1); i++ {
			foundOnce[s1[i]] = true
		}
		var foundTwice = map[byte]bool{}
		for i := 0; i < len(s2); i++ {
			if foundOnce[s2[i]] {
				foundTwice[s2[i]] = true
			}
		}
		for i := 0; i < len(s3); i++ {
			if foundTwice[s3[i]] {
				sum += toPriority(s3[i])
				break
			}
		}
	}
	fmt.Println(sum)
}
