package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type assignment struct {
	beginning int
	end       int
}

func (a assignment) contains(other assignment) bool {
	return a.beginning <= other.beginning && a.end >= other.end
}

func buildAssignment(in string) assignment {
	nums := strings.Split(in, "-")
	beginning, _ := strconv.Atoi(nums[0])
	end, _ := strconv.Atoi(nums[1])
	return assignment{beginning: beginning, end: end}
}

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "4/input.txt")
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
		assignmentInput := strings.Split(s, ",")
		assignment1 := buildAssignment(assignmentInput[0])
		assignment2 := buildAssignment(assignmentInput[1])
		if assignment1.contains(assignment2) || assignment2.contains(assignment1) {
			sum++
		}
	}
	fmt.Println(sum)
}
