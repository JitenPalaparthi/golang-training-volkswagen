package rect

func (r *Rect) Area() float32 {
	// r.area = r.Length * r.Bredth
	// return r.area
	return r.areaL()
}
func (r *Rect) areaL() float32 { //// not exportable function
	r.area = r.Length * r.Bredth
	return r.area
}
