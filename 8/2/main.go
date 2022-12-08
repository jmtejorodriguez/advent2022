package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

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

	visibility := make([][]int, height)
	for i := 0; i < height; i++ {
		visibility[i] = make([]int, width)
		for j := 1; j < width; j++ {
			visibility[i][j] = 1
		}
	}

	for i := 0; i < height; i++ {
		var view [10]int
		for j := 0; j < width; j++ {
			tree, _ := strconv.Atoi(trees[i][j : j+1])
			visibility[i][j] *= view[tree]
			for th := 0; th <= tree; th++ {
				view[th] = 1
			}
			for th := tree + 1; th < 10; th++ {
				view[th] += 1
			}
		}
	}

	for i := 0; i < height; i++ {
		var view [10]int
		for j := width - 1; j >= 0; j-- {
			tree, _ := strconv.Atoi(trees[i][j : j+1])
			visibility[i][j] *= view[tree]
			for th := 0; th <= tree; th++ {
				view[th] = 1
			}
			for th := tree + 1; th < 10; th++ {
				view[th] += 1
			}
		}
	}

	for j := 0; j < width; j++ {
		var view [10]int
		for i := 0; i < height; i++ {
			tree, _ := strconv.Atoi(trees[i][j : j+1])
			visibility[i][j] *= view[tree]
			for th := 0; th <= tree; th++ {
				view[th] = 1
			}
			for th := tree + 1; th < 10; th++ {
				view[th] += 1
			}
		}
	}

	maxV := 0
	for j := 0; j < width; j++ {
		var view [10]int
		for i := height - 1; i >= 0; i-- {
			tree, _ := strconv.Atoi(trees[i][j : j+1])
			visibility[i][j] *= view[tree]
			maxV = max(visibility[i][j], maxV)
			for th := 0; th <= tree; th++ {
				view[th] = 1
			}
			for th := tree + 1; th < 10; th++ {
				view[th] += 1
			}
		}
	}

	fmt.Println(maxV)
}
