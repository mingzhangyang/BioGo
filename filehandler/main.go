package main

import (
	gb "BioGo/filehandler/genbank"
	"log"
	"os"
)

func main() {
	gbr := gb.GBRecord{}

	if len(os.Args) < 2 {
		log.Fatal("missing argument")
	}

	err := gbr.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	println(gbr.Locus)
}
