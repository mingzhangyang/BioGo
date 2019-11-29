package genbank

import (
	"os"
	"bufio"
	"strings"
)

type holder struct {
	name string
	data []string
}

// Parse a .gb file into a GBRecord struct
func (gbr *GBRecord) Parse(fp string) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	cur := &holder{
		data: make([]string, 0, 1024),
	}
	curB := &holder {
		data: make([]string, 0, 1024),
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		head, body := string(line[:12]), string(line[12:])
		head = strings.TrimRight(head, " ")
		body = strings.TrimRight(body, " ")

		if cur.name != "FEATUREs" {
			if head == "" || head[0] == ' ' {
				cur.data = append(cur.data, body)
				continue
			}
			switch cur.name {
			case "":
				// do nothing
			case "LOCUS":
				gbr.Locus = cur.data[0]
			case "DEFINITION":
				gbr.Definition = cur.data[0]
			case "ACCESSION":
				gbr.Accession = make([]string, 0)
				for i := 0; i < len(cur.data); i++ {
					acs := strings.Split(cur.data[i], " ")
					gbr.Accession = append(gbr.Accession, acs...)
				}
			case "VERSION":
				gbr.Version = cur.data[0]
			case "DBLINK":
				gbr.Dblink = make([]string, 0, len(cur.data))
				copy(gbr.Dblink, cur.data)
			case "KEYWORDS":
				gbr.Keywords = cur.data[0]
			case "SOURCE":
				gbr.Source = newSource(cur.data)
			case "REFERENCE":
				if gbr.Reference == nil {
					gbr.Reference = make([]Reference, 0, 2)
				}
				gbr.Reference = append(gbr.Reference, newReference(cur.data))
			
			case "COMMENT":
				p, q, n := 0, 0, len(cur.data)
				for i := 0; i < n; i++ {
					if cur.data[i] == "##Genome-Annotation-Data-START##" {
						p = i
					}
					if cur.data[i] == "##Genome-Annotation-Data-END##" {
						q = i
						break
					}
				}

				if p == 0 {
					gbr.Comment = strings.Join(cur.data, " ")
					break
				}

				if p != 0 {
					gbr.Comment = strings.Join(cur.data[:p], " ")
					if q + 1 < n {
						gbr.Comment += strings.Join(cur.data[q+1:], " ")
					}
					gbr.Annotation = extracAnnotation(cur.data[p+1:q])
				}

			default:
				println("uncaught: ", head, body)
			}
	
			cur.name = head
			cur.data = cur.data[:1]
			cur.data[0] = body

			continue
		}
		
		head, body = string(line[:21]), string(line[21:])
		head = strings.TrimRight(head, " ")
		body = strings.TrimRight(body, " ")

		if head[0] != ' ' {
			cur.name = head
			cur.data = cur.data[:1]
			cur.data[0] = body

			continue
		}

		head = strings.TrimLeft(head, " ") 

		if head == "" {
			curB.data = append(curB.data, body)
			continue;
		}

		switch curB.name {
		case "":
			gbr.Features = Features{
				Genes: make([]Gene, 1024 * 8),
			}
		case "source":
			gbr.Features.Description = newFeatureDescription(curB.data)
		default:
			gbr.Features.Genes = append(gbr.Features.Genes, newGene(curB))
		}

		curB.name = head
		curB.data = curB.data[:1]
		curB.data[0] = body
	}

	switch curB.name {
	case "source":
		gbr.Features.Description = newFeatureDescription(curB.data)
	default:
		gbr.Features.Genes = append(gbr.Features.Genes, newGene(curB))
	}

	switch cur.name {

	}

	return nil
}