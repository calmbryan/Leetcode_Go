/*
Package template implements data-driven templates for generating textual
output such as HTML.
....
*/
package main

import (
	array "Leetcode_Go/array"
	"fmt"
)

func main() {
	nums := []int{2, 3, 5}
	res := array.CombinationSum(nums, 8) //[2 2 2 2] [2 3 3] [3 5]]
	fmt.Println(res)
}

func main3() {
	i := 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()
	a := []int{1, 2, 3}
	s1 := a
	s2 := a[:]
	fmt.Print(a)
	fmt.Print(s1)
	fmt.Print(s2)
	fmt.Println()
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)
	fmt.Printf("%p\n", &a)

	s := []int{1, 2, 3}
	b := [3]int{4, 5, 6}
	ref := b[:]
	fmt.Println("Existing array:\t", ref)
	t := append(s, ref...)
	fmt.Println("Existing slice:\t", s)
	fmt.Println("New slice:\t", t)
	ref[1] = -1
	fmt.Println("New slice 2:\t", t)
	s = append(s, ref...)
	fmt.Println("Existing slice:\t", s)
	s = append(s, s...)

	fmt.Println("s+s:\t\t", s)
	fmt.Println("ss" == "ss")

}
