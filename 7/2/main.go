package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

type directory struct {
	name   string
	size   int
	parent *directory
}

func (d *directory) addFile(size int) {
	d.size += size
	if d.parent != nil {
		d.parent.addFile(size)
	}
}

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "7/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	lines := bytes.Split(dat, []byte("\n"))
	directories := map[string]*directory{}
	currentDirectory := &directory{name: "/", size: 0, parent: nil}
	directories["/"] = currentDirectory

	for _, s := range lines {
		if len(s) == 0 {
			continue
		}
		if bytes.HasPrefix(s, []byte("$ cd ..")) {
			currentDirectory = currentDirectory.parent
		} else if bytes.HasPrefix(s, []byte("$ cd /")) {
			currentDirectory = directories["/"]
		} else if bytes.HasPrefix(s, []byte("$ cd")) {
			childName := currentDirectory.name + "/" + string(s[5:])
			currentDirectory = directories[childName]
		} else if bytes.HasPrefix(s, []byte("dir")) {
			childName := currentDirectory.name + "/" + string(s[4:])
			if directories[childName] == nil {
				directories[childName] = &directory{name: childName, size: 0, parent: currentDirectory}
			}
		} else if bytes.HasPrefix(s, []byte("ls")) {
			continue
		} else {
			fileSize, _ := strconv.Atoi(string(s[:bytes.Index(s, []byte(" "))]))
			currentDirectory.addFile(fileSize)
		}
	}

	candidateSize := math.MaxInt
	requiredSpace := directories["/"].size - 70000000 + 30000000

	for _, d := range directories {
		if d.size >= requiredSpace && d.size < candidateSize {
			candidateSize = d.size
		}
	}

	fmt.Println(candidateSize)
}
