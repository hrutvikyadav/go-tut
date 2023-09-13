package main

import (
	"fmt"
	"math"
)

// Interfaces are a set of method signatures
// A type can implement an Interface; i.e. that type can call the methods in the Interface

// A value of Interface type can hold any type that implements the Interface methods
// If a type does not implement an Interface or its methods, the methods are propably not defined with the concerned recv type

type IAbser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	if v == nil {
		// handle null pointer
		return 69.420
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

var a IAbser
var f MyFloat = (-math.Sqrt2)
var v Vertex = Vertex{3, 4}

// a = f  // a MyFloat implements Abser NOTE a will be a tuple (value, type) calling a.someMethod() results in value.someMethod()
// a = &v // a *Vertex implements Abser
// INFO a val of type interface can hold an underlying type that is nil and calling interface method on nil pointer could cause exeption in other languages but is handled in go by
// adding if cond inside interface method implementation To check for nil receiver;; see Abs method for *Vertex;; WARN although underlying val is nil, interface val does not become nil itself
// uninitialized Interfaces hold nil values as both the type as well as val is nil
func describeInterfaceTuple(i IAbser) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// var i Abser
// describeInterfaceTuple(i) will print the nil interface
// WARN i.Abs() is a runtime error
//
//// In the following line, v is a Vertex (not *Vertex)
//// and does NOT implement Abser.
//a = v
// NOTE to implement an interface method for a type that does not already implement it,
// just write a recv method for that recv type, No `impl` keyword exists
// INFO this is called implicit Interfaces and will decouple the DEFINITION Of an interface from its IMPLEMENTATION, which could then appear in any package

// EMPTY Interfaces i.e. Interfaces with 0 methods can hold vals of ANY type
// this is useful for functions that take in unknown types as they can now take interface{} type

// Stringer interface defined in fmt package can be implemeted on our types to enable fmt.Println-ing our type
func (v *Vertex) String() string {
	return fmt.Sprintf("Vertex\nX cordinate = %v ; Y cordinate = %v", v.X, v.Y)
}
