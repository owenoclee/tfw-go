package geo

type Vector struct {
	X int
	Y int
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector) SetX(x int) Vector {
	v.X = x
	return v
}

func (v Vector) SetY(y int) Vector {
	v.Y = y
	return v
}
