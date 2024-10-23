package square

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimter() float32 {
	return float32(4 * s)
}
