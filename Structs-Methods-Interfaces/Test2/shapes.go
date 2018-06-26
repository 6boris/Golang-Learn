package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(retangle Rectangle) float64 {
	return 2 * (retangle.Width + retangle.Height)
}
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
