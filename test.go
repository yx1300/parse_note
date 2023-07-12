package main

import "fmt"

func add[T int | float32 | float64](a, b T) T {
	return a + b
}

// aaa
// asadsa
// let me
// todo(jimmy): I'll do it later
// FIXME(tom): I'll fix it later
// fixme(jimmy): I'll fix it later
// TODO(tom): I'l do it later
// fixme(jimmy): I'll fix it later
// FIXME(jimmy): I'll fix it later
// todo(yang): I'll do it later
// FIXME(xiao): I'll fix it later
// fixme(Jeny): I'll fix it later
// TODO(jack): I'l do it later

// fixme(haha): I'll fix it later
// TODO(yang): I'l do it later
//aaaa
type mySlice[T int | float32] []T

// TODO(tom): I'l do it later
func aaa() {
	//intSlice := mySlice[int]{}
	//fmt.Println(add(1, 2))
	// todo(jimmy): I'll do it later
	var a, b uint = 2, 3
	fmt.Println(b - a)

}
