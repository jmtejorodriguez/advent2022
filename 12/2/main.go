package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

type position struct {
	x     int
	y     int
	steps int
}

func (p *position) key() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

func (p *position) equals(other *position) bool {
	return p.x == other.x && p.y == other.y
}

type node struct {
	pos  *position
	next *node
}

type queue struct {
	first *node
	last  *node
}

func (q *queue) enque(pos *position) {
	n := &node{pos: pos}
	if q.first == nil {
		q.first = n
		q.last = n
	} else {
		q.last.next = n
		q.last = n
	}
}

func (q *queue) deque() *position {
	p := q.first.pos
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	return p
}

func (q *queue) isEmpty() bool {
	return q.first == nil
}

func shouldVisit(heights [][]byte, cx int, cy int, x int, y int, visited map[string]bool) bool {
	maxx := len(heights) - 1
	maxy := len(heights[0])
	if x < 0 || x >= maxx || y < 0 || y >= maxy {
		return false
	}
	if visited[fmt.Sprintf("%d:%d", x, y)] {
		return false
	}
	ch := heights[cx][cy]
	if ch == 'S' {
		ch = 'a'
	}
	dh := heights[x][y]
	if dh == 'E' {
		dh = 'z'
	}
	if dh > ch && dh-ch > 1 {
		return false
	}
	return true
}

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "12/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	heights := bytes.Split(dat, []byte("\n"))

	finish := &position{}
	q := queue{}
	visited := make(map[string]bool)

	for i := 0; i < len(heights)-1; i++ {
		for j := 0; j < len(heights[0]); j++ {
			if heights[i][j] == 'S' || heights[i][j] == 'a' {
				q.enque(&position{x: i, y: j, steps: 0})
				visited[q.last.pos.key()] = true
			} else if heights[i][j] == 'E' {
				finish.x = i
				finish.y = j
			}
		}
	}

	for {
		current := q.deque()
		if current.equals(finish) {
			finish.steps = current.steps
			break
		}
		if shouldVisit(heights, current.x, current.y, current.x-1, current.y, visited) {
			q.enque(&position{x: current.x - 1, y: current.y, steps: current.steps + 1})
		}
		if shouldVisit(heights, current.x, current.y, current.x+1, current.y, visited) {
			q.enque(&position{x: current.x + 1, y: current.y, steps: current.steps + 1})
			visited[q.last.pos.key()] = true
		}
		if shouldVisit(heights, current.x, current.y, current.x, current.y-1, visited) {
			q.enque(&position{x: current.x, y: current.y - 1, steps: current.steps + 1})
			visited[q.last.pos.key()] = true
		}
		if shouldVisit(heights, current.x, current.y, current.x, current.y+1, visited) {
			q.enque(&position{x: current.x, y: current.y + 1, steps: current.steps + 1})
			visited[q.last.pos.key()] = true
		}
	}

	fmt.Println(finish.steps)
}
