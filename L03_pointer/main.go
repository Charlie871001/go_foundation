package main

import "fmt"

func main() {
	// var p1 *int
	// var p2 *string
	// i := 2
	// s := "hello"

	// p1 = &i
	// p2 = &s

	// p3 := &p2
	// fmt.Println(p1, p2, p3)

	// var p1 *int
	// i := 2
	// p1 = &i
	// fmt.Println(*p1 == i)
	// *p1 = 3
	// fmt.Println(i)
	a := 2
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p, &a)

	pp := &p
	fmt.Println(pp, p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)
}
