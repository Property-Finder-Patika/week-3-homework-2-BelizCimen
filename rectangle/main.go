package main

import "fmt"

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r Rectangle) Circumference() float32 {
	return 2 * (r.length + r.width)
}

func Create(l float32, w float32) (Rectangle, error) {
	if l <= 0 || w <= 0 {
		panic("width and length must be bigger than zero")
	}
	return Rectangle{length: l, width: w}, nil
}
func main() {
	var l, w float32
	fmt.Println("Enter the width of rectangle ")
	fmt.Scanln(&l)
	fmt.Println("Enter the length of rectangle ")
	fmt.Scanln(&w)
	r, _ := Create(l, w)
	fmt.Println(r)
	fmt.Println("Area of Rectangle", r.Area())
	fmt.Println("Circumference of Rectangle", r.Circumference())

}
