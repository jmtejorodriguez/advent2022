package main

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type item struct {
	element  int
	elements *[]*item
}

func (i *item) isList() bool {
	return i.elements != nil
}

func processItem(line string, currPos int) (*item, int) {
	if line[currPos] == '[' {
		els := make([]*item, 0)
		var nextPos int
		var nextItem *item
		for nextPos = currPos + 1; line[nextPos] != ']'; {
			nextItem, nextPos = processItem(line, nextPos)
			if line[nextPos] == ',' {
				nextPos++
			}
			els = append(els, nextItem)
		}
		i := &item{elements: &els}
		return i, nextPos + 1
	} else {
		var end int
		for end = currPos + 1; line[end] != ',' && line[end] != ']'; end++ {

		}
		el, _ := strconv.Atoi(line[currPos:end])
		return &item{element: el}, end
	}
}

func (i *item) compare(other *item) int {
	if !i.isList() && !other.isList() {
		return i.element - other.element
	}
	il := i
	if !il.isList() {
		il = &item{elements: &[]*item{{element: il.element}}}
	}
	ilSize := len(*il.elements)
	ol := other
	if !ol.isList() {
		ol = &item{elements: &[]*item{{element: ol.element}}}
	}
	olSize := len(*ol.elements)

	for i := 0; i < ilSize && i < olSize; i++ {
		compareItem := (*il.elements)[i].compare((*ol.elements)[i])
		if compareItem != 0 {
			return compareItem
		}
	}
	return ilSize - olSize
}

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "13/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	lines := strings.Split(string(dat), "\n")

	items := make([]*item, 0)
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		it, _ := processItem(lines[i], 0)
		items = append(items, it)
	}

	firstDivider := &item{elements: &[]*item{{elements: &[]*item{{element: 2}}}}}
	secondDivider := &item{elements: &[]*item{{elements: &[]*item{{element: 6}}}}}

	items = append(items, firstDivider, secondDivider)
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].compare(items[j]) < 0
	})

	result := 1
	for i := 0; i < len(items); i++ {
		if items[i] == firstDivider || items[i] == secondDivider {
			result *= i + 1
		}
	}
	println(result)
}
