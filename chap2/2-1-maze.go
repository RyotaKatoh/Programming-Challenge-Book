package main

import "fmt"

var maze [][]string = [][]string{
	{"#", "S", "#", "#", "#", "#", "#", "#", ".", "#"},
	{".", ".", ".", ".", ".", ".", "#", ".", ".", "#"},
	{".", "#", ".", "#", "#", ".", "#", "#", ".", "#"},
	{".", "#", ".", ".", ".", ".", ".", ".", ".", "."},
	{"#", "#", ".", "#", "#", ".", "#", "#", "#", "#"},
	{".", ".", ".", ".", "#", ".", ".", ".", ".", "#"},
	{".", "#", "#", "#", "#", "#", "#", "#", ".", "#"},
	{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
	{".", "#", "#", "#", "#", ".", "#", "#", "#", "."},
	{".", ".", ".", ".", "#", ".", ".", ".", "G", "#"},
}

type Position struct {
	X        int
	Y        int
	Progress int
}

type Node struct {
	Next *Node
	Data *Position
}

type Queue struct {
	First *Node
	Last  *Node
}

func (q *Queue) Add(n *Node) {
	if q.Last != nil {
		q.Last.Next = n
		q.Last = n
	} else {
		q.Last = n
		if q.First == nil {
			q.First = n
		}
	}
}

func (q *Queue) Remove() {
	if q.First != nil {
		q.First = q.First.Next
		if q.First == nil {
			q.Last = nil
		}
	}

	if q.First != nil {
		n := q.First
		for {
			if n.Next == nil {
				break
			}
			n = n.Next
		}
	}

}

func (q *Queue) Peek() *Node {
	return q.First
}

func (q *Queue) IsEmpty() bool {
	if q.First == nil {
		return true
	}
	return false
}

func solveMaze() {
	n := &Node{
		Data: &Position{
			X: 1,
			Y: 0,
		},
	}

	q := &Queue{}
	q.Add(n)

	lenY := len(maze)
	lenX := len(maze[0])
	posX := n.Data.X
	posY := n.Data.Y

	moves := []*Position{
		&Position{X: -1, Y: 0},
		&Position{X: 0, Y: -1},
		&Position{X: 1, Y: 0},
		&Position{X: 0, Y: 1},
	}

	for {
		if q.IsEmpty() {
			fmt.Println("NO GOAL")
			return
		}

		current := q.Peek()
		posX = current.Data.X
		posY = current.Data.Y
		q.Remove()

		for _, m := range moves {
			if posY+m.Y >= 0 && posY+m.Y < lenY && posX+m.X >= 0 && posX+m.X < lenX && maze[posY+m.Y][posX+m.X] != "#" {
				newN := &Node{
					Data: &Position{
						X:        posX + m.X,
						Y:        posY + m.Y,
						Progress: current.Data.Progress + 1,
					},
				}
				q.Add(newN)

				if maze[newN.Data.Y][newN.Data.X] == "G" {
					fmt.Println("GOAL!!!")
					fmt.Println("ans: ", newN.Data.Progress)
					return
				}

				maze[posY+m.Y][posX+m.X] = "#"
			}
		}

	}

}

func main() {
	solveMaze()
}
