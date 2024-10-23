package rect

func (r *Rect) Perimeter() float32 {
	return r.perimeterL() // can be called locally
}

func (r *Rect) perimeterL() float32 { // not exportable function
	r.perimeter = 2 * (r.Length + r.Bredth)
	return r.perimeter
}
