package square

import "errors"

// type Square float32

// func (s Square) Area() float32 {
// 	return float32(s * s)
// }

//	func (s Square) Perimter() float32 {
//		return float32(4 * s)
//	}
var Side float32 // since Side S is Upper case , it is exported

type Square struct{} // empty struct

func New(side float32) (*Square, error) {
	if side == 0 {
		return nil, errors.New("invalid input")
	}
	Side = side
	return &Square{}, nil
}
func (s Square) Area() float32 {
	return Side * Side
}

func (s Square) Perimter() float32 {
	return 4 * Side
}

// Side global variable is for demo just to tell that , can also use like this
