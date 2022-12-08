package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "8/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	trees := strings.Split(string(dat), "\n")

	width := len(trees[0])
	height := len(trees) - 1

	visible := make([][]bool, height)
	for i := 0; i < height; i++ {
		visible[i] = make([]bool, width)
	}

	for i := 0; i < height; i++ {
		visible[i][0] = true
		max := trees[i][0]
		for j := 1; j < width; j++ {
			if trees[i][j] > max {
				visible[i][j] = true
				max = trees[i][j]
			}
		}
		visible[i][width-1] = true
		max = trees[i][width-1]
		for j := width - 2; j >= 0; j-- {
			if trees[i][j] > max {
				visible[i][j] = true
				max = trees[i][j]
			}
		}
	}

	for j := 0; j < width; j++ {
		visible[0][j] = true
		max := trees[0][j]
		for i := 1; i < height; i++ {
			if trees[i][j] > max {
				visible[i][j] = true
				max = trees[i][j]
			}
		}
		visible[height-1][j] = true
		max = trees[height-1][j]
		for i := height - 2; i >= 0; i-- {
			if trees[i][j] > max {
				visible[i][j] = true
				max = trees[i][j]
			}
		}
	}

	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if visible[i][j] {
				count++
			}
		}
	}

	fmt.Println(count)
}
