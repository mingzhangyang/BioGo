package main

import (
	gb "BioGo/filehandler/genbank"
	"fmt"
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
	fmt.Println(gbr.Features.Description)
	fmt.Println(gbr.DbLink)
	fmt.Println(gbr.Accession)
	fmt.Println(gbr.Version)
	fmt.Println(gbr.Contig)
	fmt.Println(gbr.Comment)
	for k, v := range gbr.Annotation {
		fmt.Println(k, "---", v)
	}
	fmt.Println(len(gbr.Features.Genes))
	fmt.Println(gbr.Features.Genes[0])
	fmt.Println(gbr.Features.Genes[1])
	fmt.Println(gbr.Features.Genes[2])
}
