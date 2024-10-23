package cuboid

type Cuboid struct {
	A, P, H float32 // Area and Perimeter two fields
	L, B    float32
}

// methods
// methods contain receivers
func (r *Cuboid) Area() float32 {
	r.A = r.L * r.B * r.H
	return r.A
}

func (r *Cuboid) Perimeter() float32 {
	r.P = 2 * (r.L + r.B + r.H)
	return r.P
}
func (r *Cuboid) GetType() string {
	return "Cuboid"
}
