package main

import (
	"math"
)

// Create a Circle struct type
type Circle struct {
	Radius float64
	Area   float64
}

// Create a rectangle struct type
type Rect struct {
	Width  float64
	Height float64
	Area   float64
}

// Function to create new Circle
func createCircle(radius float64) Circle {
	c := Circle{
		Radius: radius,
		Area:   math.Pi * radius * radius,
	}
	// return the new circle
	return c
}

// Function to create new Rectangle
func createRetangle(width float64, height float64) Rect {
	r := Rect{
		Width:  width,
		Height: height,
		Area:   width * height,
	}
	// return the new rectangle
	return r
}

// Fuction that takes a type struct and updates radius of a circle and calculates a new area
func (c *Circle) updateCircle(radius float64) {
	c.Radius = radius
	c.Area = math.Pi * radius * radius
}

// Function that takes a type struct and updates width and height of a rectangle and calculates a new area
func (r *Rect) updateRectangle(width float64, height float64) {
	r.Width = width
	r.Height = height
	r.Area = width * height
}

// Function to calculate circumference for circle
func (c *Circle) circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// Function to calculate circumference of rectangle
func (r *Rect) circumference() float64 {
	return 2*r.Width + 2*r.Height
}

// NOTE: Both circle and rectangle have circumference() function with SIMILAR behaviour, but different implementation.
// So both circle and rectangle can be grouped into a INTERFACE type with a SIMILAR circumference() behaviour
// INTERFACES lets you look at the above circle and rect types as if they have SIMILAR behaviour thereby abstracting other DISSIMILAR details
// They can BOTH have area. They can BOTH have circumference.

// CREATING AN INTERFACE FOR CIRCUMFERENCE
// So an interface groups types together based on their methods
type shape interface {
	circumference() float64
}

// Function that takes in a shape and exectues a code
func calculateCircumference(s shape) float64 {
	//fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumference())
	return s.circumference()
}
