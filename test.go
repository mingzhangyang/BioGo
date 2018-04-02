package main

import (
	"fmt"

	"./sequence"
)

func main() {
	s, _ := sequence.NewDNA("agtcgatcgtagcta")
	fmt.Println(s)
	fmt.Println(s.Range(5, 10))
	// fmt.Println(s.Uniq())
	// fmt.Println(s.Length())
	// fmt.Println(s.Composition())
	// fmt.Println(s.GCContent())
	// fmt.Println(s.Complement())
	// fmt.Println(s.ReverseComplement())
}
