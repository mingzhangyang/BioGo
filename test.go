package main

import (
	// "log"
	"BioGo/filehandler"
	"os"
	"fmt"
	// "./machinary"
	// "./sequence"
	// "./utilities"
	// ds "./datastructure"
	// m "./machinary"
	// "BioGo/util"

)

func main() {
	// fq := filehandler.Fastq{"./sampledata/SRR835775_1.first1000.fastq", "test"}
	// fmt.Println(fq.Hist())
	// m := fq.QCounts()
	// s := make([]int, 60)
	// for k, v := range m {
	// 	s[k] = v
	// }
	// fmt.Println(s)
	// s := `AACTAGCACTAGCTGTTGCTATCGTACGTAGTTCATTGGTCATCGACCGGGTCATGCATCTAGCATCGTAGCATGCTAGCGATCTAGCTAGTCGTAGCTAGTCAGCGTAGCGTACGTAGCTAGCTAGCTAGTCGATCGATGCTAGCTAGTCGTAGCTAGGTTCTATGCT`

	// d1 := util.CountSubsequentOneChar(s, nil)
	// d2 := util.CountPrecedingOneChar(s, nil)

	// fmt.Println(d1)
	// fmt.Println(d2)

	// fmt.Println(m.I)
	// fmt.Println(m.II)
	// fmt.Println(m.IV)
	// fmt.Println(m.IX)
	// s := ds.Stack{}
	// s.Push("A")
	// s.Push("B")
	// s.Push("C")
	// s.Push("D")
	// fmt.Println(s)
	// fmt.Println(s.Length())
	// s.Pop()
	// s.Pop()
	// s.Pop()
	// fmt.Println(s)
	// fmt.Println(s.Length())

	// fmt.Println("***********************************")

	// q := ds.Queue{}
	// q.EnQueue("X")
	// q.EnQueue("Y")
	// q.EnQueue("Z")
	// fmt.Println(q)
	// fmt.Println(q.Length())
	// q.DeQueue()
	// q.DeQueue()
	// fmt.Println(q)
	// fmt.Println(q.Length())
	// c := utilities.NewCircularIndex(0, 10)
	// c.SetStep(6)
	// fmt.Println(c.Next())
	// fmt.Println(c.Next())
	// fmt.Println(c.Next())
	// fmt.Println(utilities.HammingDistance("abcdefg", "adcegfb"))
	// s, _ := sequence.NewDNA("agtcgatcgtaggatccta")
	// fmt.Println(s)
	// fmt.Println(s.Range(-5))
	// e, _ := machinary.NewRE_EnzymeFromName("BamHI")
	// fmt.Println(e.SearchDNA(s))
	// fmt.Println(s.Uniq())
	// fmt.Println(s.Length())
	// fmt.Println(s.Composition())
	// fmt.Println(s.GCContent())
	// fmt.Println(s.Complement())
	// fmt.Println(s.ReverseComplement())

	// fa, err := filehandler.ReadFasta("./sampledata/multiple.fasta")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(*fa)

	// s := `MNGYISNANYSVKKCIFLLFINHRLVESAALRKAIETVYAAYLPKNTHPFLYLSLEISPQNVDVNVHPTK
	// HEVHFLHEESILQRVQQHIESKLLGSNSSRMYFTQTLLPGLAGPSGEAARPTTGVASSSTSGSGDKVYAY
	// QMVRTDSREQKLDAFLQPVSSLGPSQPQDPAPVRGARTEGSPERATREDEEMLALPAPAEAAAESENLER
	// ESLMETSDAAQKAAPTSSPGSSRKRHREDSDVEMVENASGKEMTAACYPRRRIINLTSVLSLQEEISERC
	// HETLREMLRNHSFVGCVNPQWALAQHQTKLYLLNTTKLSEELFYQILIYDFANFGVLRLSEPAPLFDLAM
	// LALDSPESGWTEDDGPKEGLAEYIVEFLKKKAEMLADYFSVEIDEEGNLIGLPLLIDSYVPPLEGLPIFI
	// LRLATEVNWDEEKECFESLSKECAMFYSIRKQYILEESTLSGQQSDMPGSTSKPWKWTVEHIIYKAFRSH
	// LLPPKHFTEDGNVLQLANLPDLYKVFERC`

	// wc := util.NewWordCounter(s, 3, "ARNDCEQGHILKMFPSTWYV")
	// fmt.Println(wc.Ratio())
	// fmt.Println(wc)
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("missing arguments")
	} else {
		c := make(chan struct{}, 8)
		for i := 1; i < len(os.Args); i++ {
			go func(f string) {
				var gr filehandler.GFF3Reader
				err := gr.Convert2JSON(f)
				if err != nil {
					fmt.Printf("failed to convert file %s for the reason that %s\n", f, err)
				}
				fmt.Printf("file %s is converted to JSON successfully\n", f)
				c <- struct{}{}
			}(os.Args[i])
		}
		N := len(os.Args) - 1
		for i := 0; i < N; i++ {
			<-c
		}
	}

}
