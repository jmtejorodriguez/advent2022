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
	if p.x+2 == head.x && p.y < head.y {
		p.x++
		p.y++
	} else if p.x+2 == head.x && p.y > head.y {
		p.x++
		p.y--
	} else if p.x+2 == head.x {
		p.x++
	} else if p.x-2 == head.x && p.y < head.y {
		p.x--
		p.y++
	} else if p.x-2 == head.x && p.y > head.y {
		p.x--
		p.y--
	} else if p.x-2 == head.x {
		p.x--
	} else if p.y+2 == head.y && p.x < head.x {
		p.y++
		p.x++
	} else if p.y+2 == head.y && p.x > head.x {
		p.y++
		p.x--
	} else if p.y+2 == head.y {
		p.y++
	} else if p.y-2 == head.y && p.x < head.x {
		p.y--
		p.x++
	} else if p.y-2 == head.y && p.x > head.x {
		p.y--
		p.x--
	} else if p.y-2 == head.y {
		p.y--
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

	rope := make([]position, 10)
	for i := 0; i < 10; i++ {
		rope[i] = position{}
	}

	visited := map[string]bool{}
	visited[rope[9].key()] = true
	for _, movement := range movements {
		if len(movement) == 0 {
			continue
		}
		direction := movement[0]
		times, _ := strconv.Atoi(string(movement[2:]))
		for i := 0; i < times; i++ {
			rope[0].move(direction)
			for j := 1; j < 10; j++ {
				rope[j].follow(rope[j-1])
			}
			visited[rope[9].key()] = true
		}
	}

	fmt.Println(len(visited))
}
