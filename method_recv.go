package main

import (
// "math"
)

// Define methods on TYPES with method Receivers
// This is just a func + a receiver TYPE? argument between func keyword and method name
// NOTE: and this means that the receiver TYPE can call this method

// WARN: both type and recv method must be in same package
type Vertex struct {
	X, Y float64
}

// recv can be passed by value or by reference
// value here is just syntactic sugar for pointers under the hood
//func (v Vertex) Abs() float64 {
//return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
// INFO: above recv fn can be written as a normal fn but it will not be restricted to some receiver type
// but rather can be called on any type of vertex

//func Abs(v Vertex) float64 {
//return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

func (vp *Vertex) Scale(f float64) {
	// scale vertex by factor f
	vp.X *= f
	vp.Y *= f
}

// NOTE methods with p recv called by val recv will work fine; but normal funcs need to be called with the exact specified param - either T or *T
