package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	b.shapes = append(b.shapes, shape)
	if len(b.shapes) > b.shapesCapacity {
		return fmt.Errorf("box capacity has been exceeded")
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {

	if i > b.shapesCapacity {
		return nil, fmt.Errorf("index is out of range")
	}

	if b.shapes[i] == nil {
		return nil, fmt.Errorf("shape doesn't exist")
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i > b.shapesCapacity {
		return nil, fmt.Errorf("index is out of range")
	}

	if b.shapes[i] == nil {
		return nil, fmt.Errorf("shape doesn't exist")
	}

	tempShape := b.shapes[i]
	b.shapes[i] = nil

	return tempShape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i > b.shapesCapacity {
		return nil, fmt.Errorf("index is out of range")
	}

	if b.shapes[i] == nil {
		return nil, fmt.Errorf("shape doesn't exist")
	}

	tempShape := b.shapes[i]
	b.shapes[i] = shape

	return tempShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {

	var sum float64

	for _, shape := range b.shapes {

		sum += shape.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {

	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}

	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {

	var trianglesExist bool = false

	for _, shape := range b.shapes {
		_, ok := shape.(Triangle)

		if ok {
			trianglesExist = true
			shape = nil
		}
	}

	if !trianglesExist {
		return fmt.Errorf("There are no triangles in the box")
	}

	return nil
}
