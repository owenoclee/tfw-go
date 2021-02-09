package main

type vector struct {
	x int
	y int
}

func (v1 vector) add(v2 vector) vector {
	return vector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

type rect struct {
	topLeft     vector
	bottomRight vector
}

func (r rect) HorizontalCells() int {
	return r.bottomRight.x - r.topLeft.x + 1
}

func (r rect) VerticalCells() int {
	return r.bottomRight.y - r.topLeft.y + 1
}

func (r rect) IsValid() bool {
	return r.topLeft.x <= r.bottomRight.x && r.topLeft.y <= r.bottomRight.y
}

func (r rect) VectorInBounds(v vector) bool {
	return v.x >= r.topLeft.x &&
		v.x <= r.bottomRight.x &&
		v.y >= r.topLeft.y &&
		v.y <= r.bottomRight.y
}
