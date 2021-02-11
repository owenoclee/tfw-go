package geom

type Rect struct {
	TopLeft     Vector
	BottomRight Vector
}

func (r Rect) HorizontalCells() int {
	return r.BottomRight.X - r.TopLeft.X + 1
}

func (r Rect) VerticalCells() int {
	return r.BottomRight.Y - r.TopLeft.Y + 1
}

func (r Rect) IsValid() bool {
	return r.TopLeft.X <= r.BottomRight.X && r.TopLeft.Y <= r.BottomRight.Y
}

func (r Rect) VectorInBounds(v Vector) bool {
	return v.X >= r.TopLeft.X &&
		v.X <= r.BottomRight.X &&
		v.Y >= r.TopLeft.Y &&
		v.Y <= r.BottomRight.Y
}

func (r Rect) RectInBounds(r2 Rect) bool {
	return r.VectorInBounds(r2.TopLeft) && r.VectorInBounds(r2.BottomRight)
}

func (r Rect) Transform(v Vector) Rect {
	return Rect{
		TopLeft:     r.TopLeft.Add(v),
		BottomRight: r.BottomRight.Add(v),
	}
}
