package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type position struct {
	x int
	y int
}

func (p *position) move(direction byte) {
	switch direction {
	case 'U':
		p.y--
	case 'D':
		p.y++
	case 'L':
		p.x--
	case 'R':
		p.x++
	}
}

func (p *position) follow(head position) {
	if p.x+2 == head.x {
		p.x++
		p.y = head.y
	} else if p.x-2 == head.x {
		p.x--
		p.y = head.y
	} else if p.y+2 == head.y {
		p.y++
		p.x = head.x
	} else if p.y-2 == head.y {
		p.y--
		p.x = head.x
	}
}

func (p *position) key() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "9/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	movements := bytes.Split(dat, []byte("\n"))

	hpos, tpos := position{}, position{}

	visited := map[string]bool{}
	visited[tpos.key()] = true
	for _, movement := range movements {
		if len(movement) == 0 {
			continue
		}
		direction := movement[0]
		times, _ := strconv.Atoi(string(movement[2:]))
		for i := 0; i < times; i++ {
			hpos.move(direction)
			tpos.follow(hpos)
			visited[tpos.key()] = true
		}
	}

	fmt.Println(len(visited))
}
