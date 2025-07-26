package rectangle

type Shaper interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}