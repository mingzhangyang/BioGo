package genbank

import (
	"os"
	"bufio"
	"strings"
	"path/filepath"
	"errors"
)

type holder struct {
	name string
	data []string
}

// Parse a .gb file into a GBRecord struct
func (gbr *GBRecord) Parse(fp string) error {
	ext := filepath.Ext(fp)
	if ext != ".gb" {
		return errors.New("genbank file is supposed to contain name extension .gb")
	}
	
	f, err := os.Open(fp)
	if err != nil {
		return err
	}

	base := filepath.Base(fp)
	gbr.Name = strings.Split(base, ext)[0]

	scanner := bufio.NewScanner(f)
	cur := &holder{
		data: make([]string, 0, 1024),
	}
	curB := &holder {
		data: make([]string, 0, 1024),
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) < 12 {
			println("illegal line: ", string(line))
			continue
		}
		head, body := string(line[:12]), string(line[12:])
		head = strings.TrimRight(head, " ")
		body = strings.TrimRight(body, " ")

		if cur.name != "FEATURES" {
			if head == "" {
				cur.data = append(cur.data, body)
				continue
			}
			if head[0] == ' ' {
				cur.data = append(cur.data, string(line))
				continue
			}
			switch cur.name {
			case "":
				// do nothing
			case "LOCUS":
				gbr.Locus = strings.Join(cur.data, " ")
			case "DEFINITION":
				gbr.Definition = strings.Join(cur.data, " ")
			case "ACCESSION":
				gbr.Accession = make([]string, 0)
				for i := 0; i < len(cur.data); i++ {
					acs := strings.Split(cur.data[i], " ")
					gbr.Accession = append(gbr.Accession, acs...)
				}
			case "VERSION":
				gbr.Version = strings.Join(cur.data, " ")
			case "DBLINK":
				gbr.Dblink = make([]string, 0, len(cur.data))
				copy(gbr.Dblink, cur.data)
			case "KEYWORDS":
				gbr.Keywords = strings.Join(cur.data, " ")
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

			case "CONTIG":
				gbr.Contig = strings.Join(cur.data, " ")
			default:
				println("uncaught: ", head, body)
			}
	
			println("debugging.... |", head)
			cur.name = head
			cur.data = cur.data[:1]
			cur.data[0] = body

			continue
		}
		
		
		head, body = string(line[:21]), string(line[21:])
		head = strings.TrimRight(head, " ")
		body = strings.TrimRight(body, " ")

		if head == "" {
			curB.data = append(curB.data, body)
			continue;
		}
		
		if head[0] != ' ' {
			cur.name = head
			cur.data = cur.data[:1]
			cur.data[0] = body

			continue
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

	if curB.name != "" {
		println(curB.name, "|...")
		switch curB.name {
		case "source":
			gbr.Features.Description = newFeatureDescription(curB.data)
		default:
			println("debuging.....\n", curB.name)
			gbr.Features.Genes = append(gbr.Features.Genes, newGene(curB))
		}
	}

	if cur.name != "" {
		println(cur.name, "|...")
		switch cur.name {
		case "CONTIG":
			// do nothing since it has been handle before
		default:
			println("bypass lines: ", strings.Join(cur.data, " "))
		}
	}

	return nil
}