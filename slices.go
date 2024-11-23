package main

import (
	"fmt"
	"math"
)

var sliceInts []int = []int{
	1,
	2,
	3,
	4,
	5,
}

var sliceFloats []float64 = []float64{
	1.32,
	3.33,
	4.36,
	3.22,
}

var sliceStrings []string = []string{
	"Jace",
	"David",
	"Jake",
}

var sliceMaps []map[int]string = []map[int]string{
	{1: "Jace", 2: "Henry", 3: "David"},
	{1: "Joe", 2: "Joanna", 3: "Healy"},
	{1: "Emily", 2: "Donna", 3: "Donald"},
}

type structForSlice struct {
	name         string
	age          int
	relationship bool
}

var sliceStructs []structForSlice = []structForSlice{
	{name: "J", age: 32, relationship: true},
	{name: "R", age: 28, relationship: true},
	{name: "A", age: 17, relationship: false},
}

// Creating slice for Interface
type square struct {
	length float64
}

type circle struct {
	radius float64
}

// Square Methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumference() float64 {
	return s.length * 4
}

// Circle Methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

type interfaceForSlice interface {
	circumference() float64
	area() float64
}

func printShapeInfo(s interfaceForSlice) {
	fmt.Printf("area of %T is: %0.2f\n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f\n", s, s.circumference())
}

var sliceInterface []interfaceForSlice = []interfaceForSlice{
	square{length: 15.2},
	circle{radius: 7.5},
	circle{radius: 12.3},
	square{length: 4.9},
}
