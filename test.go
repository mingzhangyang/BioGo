package main

import (
	"fmt"

	"./machinary"
	"./sequence"
)

func main() {
	s, _ := sequence.NewDNA("agtcgatcgtaggatccta")
	fmt.Println(s)
	fmt.Println(s.Range(-5))
	e, _ := machinary.NewRE_EnzymeFromName("BamHI")
	fmt.Println(e.SearchDNA(s))
	// fmt.Println(s.Uniq())
	// fmt.Println(s.Length())
	// fmt.Println(s.Composition())
	// fmt.Println(s.GCContent())
	// fmt.Println(s.Complement())
	// fmt.Println(s.ReverseComplement())
}
