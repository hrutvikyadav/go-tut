package main

import (
	"fmt"
	"runtime"
)
import "rsc.io/quote/v4"

func main() {
	// defers the following function call when this function exits
	// arguments are evaluated when defer someF(arg1, arg2) is encountered in code, call is made when reached end of block
	defer fmt.Println("defer says, Exiting now")
	defer fmt.Println("Hey I am first in line, LIFO remember") // this is executed first

	fmt.Println(quote.Go())
	fmt.Println("hi from go")

	switch my_os := runtime.GOOS; my_os {
	case "darwin":
		fmt.Println("OS x")
	case "linux":
		fmt.Println("My os")
	default:
		fmt.Printf("oh god: %v", my_os)
	}

	// POINTERS
	var p *int
	i := 10
	fmt.Println("Original var, original value: ", i)
	p = &i
	*p = *p + 69
	fmt.Println("Pointer: ", p, "Original var: ", i, "Deref value after mutation: ", *p)

	// STRUCTS
	type Me struct {
		name   string
		number int
		height float32
		weight float32
	}
	var m Me
	//access fields with `.` operator
	m.name = "Hrx"
	fmt.Println(m, m.name)
	m = Me{"Hrxxx", 69, 183.45, 60.2}
	fmt.Println("NEW me:", m)

	// access fields with *struct
	psm := &m
	// struct data accessed with pointer with both notations, NOTE: different notations in below call
	info := fmt.Sprintf("With s pointer\nName: %v, Height: %v", (*psm).name, psm.height)
	fmt.Println(info)

	var _ *Me
	// struct literal returning pointer along with omitted fields assigned their zero value (0 for numeric types, ""for strings, nil for errors)
	p_pushkar := &Me{name: "Pus", number: 111}
	fmt.Println("Pushkar is: \n", p_pushkar.name, p_pushkar.number, p_pushkar.height)

	// Arrays NOTE: non resizable
	var arr [7]int
	arr[0] = 420
	arr[1] = 69
	fmt.Println(arr)

	arrMe := [10]Me{*psm, m, *p_pushkar}
	fmt.Println(arrMe)

	// Slices grow/shrink - able arrays with zero value = nil
	// It is like a reference to a slice of a bigger array, and modifying slices will modify actual data in the array to which the slice belongs
	var threeImpPeople []Me = arrMe[0:3] // include start; exclude end; when bounds are omitted, creates a sclice of the whole array
	fmt.Println(threeImpPeople, "Len: ", len(threeImpPeople), "Cap: ", cap(threeImpPeople))

	coordinates := []struct{ X, Y int }{{X: 1, Y: 2}, {X: 99, Y: 98}}
	fmt.Println(coordinates)

	// resliced people to accomodate for more imp people
	var fourImpPeople []Me = arrMe[0:4]
	fmt.Println("Newly sliced: ", fourImpPeople, len(fourImpPeople), cap(fourImpPeople))

	// invalid slice - slice index greater than array capacity
	//unImpPeople := arrMe[:11] // idx out of bounds

	// dynamic allocation of slices in this case with make
	s_ints := make([]int, 9, 69) // creates a zero valued array of ints in this case with len 9 and cap 69
	fmt.Println(s_ints)

	s_ints = s_ints[:(cap(s_ints) - 50)] // capacity ++
	fmt.Println(s_ints)

	s_ints[1] = 29
	for i := 0; i < len(s_ints); i++ {
		if s_ints[i] != 0 {
			fmt.Println(s_ints[i])
		}
	}

	s_ints[2] = 6
	s_ints[3] = 2000
	for i, v := range s_ints {
		if v != 0 {
			fmt.Println(i, v)
		}
	}

	c := WordCount("count me bitch ! bitch?")
	fmt.Println(c)

	fibValue := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibValue())
	}
	fmt.Println(fibValue()) // prints next number in series i.e. 11th one

	fibValue2 := Fibonacci() // another closure, bound to a different lastTwo variable than the above funtion `fibValue()`

	fmt.Println("Another series of 15")
	for i := 0; i < 15; i++ {
		fmt.Println(fibValue2())
	}

	vertex := Vertex{X: 20.22, Y: 45.79}
	fmt.Println(vertex.Abs())

	vp := &vertex
	vp.Scale(10) // scale must be called on a pointer recv to make sure that it modifies the type pointed to and not a copy of that type; like what happens in val recv
	fmt.Println(*vp)

	// also works and it inerpreted as (&vertex).Scale(20)
	vertex.Scale(20)

	// Interfaces
	var ia IAbser
	ia = &vertex
	// assert if val in ia is a certain type
	my_v, ok := ia.(*Vertex)
	fmt.Println(my_v, ok, ia.Abs(), my_v.Abs())

	// switch over multiple type assertions
	switch v := ia.(type) {
	case *Vertex:
		fmt.Printf("%v scaled by 69 is ", v)
		v.Scale(69)
		fmt.Println(v)
	case MyFloat:
		fmt.Println("My Float Abs is", v.Abs())
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	fmt.Println("Before Stringer interface ^^")
	fmt.Println("After Stringer interface --")
	fmt.Println("Look at where vertex reached now!!\n", my_v)


	fmt.Println(Sqrt(-2))
	fmt.Println(ErrNegativeSqrt(-10).Error())

    DemoStringReader()

    mar := MyAReader{}

    buf := make([]byte, 8)
    num, a_err := mar.Read(buf)
    fmt.Println(num, a_err, buf)
    fmt.Printf("Buffer cont: %q\n", buf[:num])


	l1 := &List[int]{&List[int]{&List[int]{&List[int]{nil, 120}, 110}, 100}, 90}

	l2 := List[string]{&List[string]{&List[string]{&List[string]{nil, "Monitor"}, "CPU"}, "Printer"}, "Speaker"}

	fmt.Println(l1.Tail())
	fmt.Println(l1.LookUp(90))

	fmt.Println(l2.Tail())
	fmt.Println(l2.LookUp("Monitor"))

	fmt.Println("Before app: ", l1)

	l1 = l1.Append(&List[int]{nil, 420})
	fmt.Println(l1.Tail())
	fmt.Println(l1.LookUp(90))

	fmt.Println("After app: ", l1)

	l3 := l2.Append(&List[string]{&List[string]{&List[string]{&List[string]{nil, "DualMonitor"}, "Cycler"}, "Frame"}, "FidgetSpinner"})

	fmt.Println(l3)
	fmt.Println(l3.Tail())
	fmt.Println(l3.LookUp("FidgetSpinner"))
	
	fmt.Println("All elems in L3:")
	l3.PrintAll()

}
