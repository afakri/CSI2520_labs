package main

import "fmt"

func main() {
	tree := PtTree{Point{2, 3},
		&PtTree{Point{5, 1},
			&PtTree{Point{2, 2}, nil, nil},
			&PtTree{Point{8, 3}, &PtTree{Point{1, 6}, nil, nil}, nil}},
		&PtTree{Point{4, 7},
			&PtTree{Point{7, 2},
				&PtTree{Point{6, 4}, nil, nil},
				&PtTree{Point{0, 9}, nil, nil}},
			&PtTree{Point{3, 6}, nil, nil}}}
	tree.postOrder()
	fmt.Println("")
	u, v := 7, 2
	if tree.find(u, v) {
		fmt.Printf("Found: %d %d \n", u, v)
	} else {
		fmt.Printf("Not Found\n")
	}
	x, y := 8, 6
	if tree.find(x, y) {
		fmt.Printf("Found: %d %d \n", u, v)
	} else {
		fmt.Printf("Not Found\n")
	}
}

// Point is ...
type Point struct {
	x int
	y int
}

//PtTree is ...
type PtTree struct {
	pt    Point
	left  *PtTree
	right *PtTree
}

func (p *PtTree) postOrder() {
	if p == nil {
		return
	}
	p.left.postOrder()
	p.right.postOrder()
	fmt.Print(p.pt)
}
func (p *PtTree) equals(x int, y int) bool {
	if p.pt.x == x && p.pt.y == y {
		return true
	}
	return false

}

func (p *PtTree) find(x int, y int) bool {
	if p == nil {
		return false
	}
	if p.equals(x, y) {
		return true
	}
	if p.left.find(x, y) {
		return true
	}
	return p.right.find(x, y)

}
