package rect

type Rect struct {
	A, P float32 // Area and Perimeter two fields
	L, B float32
}

// methods
// methods contain receivers
func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}
func (r *Rect) GetType() string {
	return "Rect"
}
