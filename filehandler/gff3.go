package filehandler

import (
	"encoding/json"
	"strconv"
	"strings"
	"os"
	"bufio"
	"fmt"
	"errors"
)

// GFF3Object format
type GFF3Object struct {
	SeqID string
	Source string
	Type string
	Start int
	End int
	Score float64
	Strand string
	Phase string
	Attributes map[string]string
}

// Load the data
func (g *GFF3Object) Load(line string) error {
	fields := strings.Split(line, "\t")
	if (len(fields) != 9) {
		return errors.New("not enough fields")
	}
	g.SeqID = fields[0]
	g.Source = fields[1]
	g.Type = fields[2]
	var i int
	var err error
	if fields[3] != "." {
		i, err = strconv.Atoi(fields[3])
		if err != nil {
			return err
		}
	}
	g.Start = i
	if fields[4] != "." {
		i, err = strconv.Atoi(fields[4])
		if err != nil {
			return err
		}
	}
	g.End = i
	if fields[5] != "." {
		f, err := strconv.ParseFloat(fields[5], 64)
		if err != nil {
			return err
		}
		g.Score = f
	}
	g.Strand = fields[6]
	g.Phase = fields[7]
	if fields[8] != "." {
		var m = make(map[string]string)
		arr := strings.Split(fields[8], ";")
		for i := 0; i < len(arr); i++ {
			kv := strings.Split(arr[i], "=")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			} else {
				m[kv[0]] = arr[i]
			}
		}
		g.Attributes = m
	}
	return nil
}


// GFF3Content of a .gff3 file
type GFF3Content struct {
	Comment map[string]string
	Content []*GFF3Object
}

// GFF3Reader read a .gff3 file
type GFF3Reader struct{}

// Read a .gff3 file
func (gr GFF3Reader) Read(file string) (*GFF3Content, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var gc = GFF3Content{
		Comment: make(map[string]string),
		Content: make([]*GFF3Object, 0),
	}
	var scanner = bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if strings.HasPrefix(line, "##") {
			line = strings.TrimPrefix(line, "##")
			a := strings.Split(line, " ")
			if len(a) == 2 {
				gc.Comment[a[0]] = a[1]
			} else {
				gc.Comment[a[0]] = line
			}
		} else {
			if line != "" {
				var g GFF3Object;
				err := g.Load(line)
				if err != nil {
					fmt.Printf("bad line: %s\n", line)
				} else {
					gc.Content = append(gc.Content, &g)
				}
			} else {
				fmt.Println("empty line ignored")
			}
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return &gc, nil
}

// Convert2JSON convert .gff3 file to JSON file
func (gr GFF3Reader) Convert2JSON(file string) error {
	gc, err := gr.Read(file)
	if err != nil {
		return err
	}
	output, err := os.Create(file + ".json")
	if err != nil {
		return err
	}
	defer output.Close()
	out := bufio.NewWriter(output)
	bs, err := json.MarshalIndent(gc, "", "    ")
	if err != nil {
		return err
	}
	_, err = out.Write(bs)
	if err != nil {
		return err
	}
	return nil
}