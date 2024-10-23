package rect

import "errors"

// var (
// 	V = 100
// )

// const (
// 	PI = 3.14
// )

var ErrInvaliInput = errors.New("invalid input")

type Rect struct { // exportable
	Length, Bredth  float32 // exportable
	area, perimeter float32 // not exportable
}

// Kind of a constructor, Idiomatic approach of constructing an object
func New(l, b float32) (*Rect, error) {
	if l == 0 || b == 0 {
		return nil, ErrInvaliInput
	}
	return &Rect{Length: l, Bredth: b}, nil
}
