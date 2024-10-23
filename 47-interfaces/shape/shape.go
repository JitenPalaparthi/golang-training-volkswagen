package shape

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}

type IShape interface {
	IArea
	IPerimeter
	GetType() string
}
