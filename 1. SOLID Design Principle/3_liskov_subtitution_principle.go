package main

import (
	"fmt"
)

// Liskov Subtitution Principle

// Sized ..
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

// Rectangle ..
type Rectangle struct {
	width, height int
}

// GetWidth ..
func (r *Rectangle) GetWidth() int {
	return r.width
}

// SetWidth ...
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

// GetHeight ..
func (r *Rectangle) GetHeight() int {
	return r.height
}

// SetHeight ...
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// Square ...
// modified LSP
// If a function takes an interface and
// works with a type T that implements this
// interface, any structure that aggregates T
// should also be usable in that function.
type Square struct {
	Rectangle
}

// NewSquare ...
func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// SetWidth Square...
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

// SetHeight Square...
func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// Square2 ...
type Square2 struct {
	size int // width, height
}

// Rectangle ...
func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

// UseIt ...
func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected an area of", expectedArea, "but got", actualArea)
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
