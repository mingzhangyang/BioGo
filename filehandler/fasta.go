package filehandler

import (
	// seq "BioGo/sequence"
	"os"
	"bufio"
)

// Fasta type represent a handler to deal with fasta files
type Fasta struct {
	Source string
	Identifiers []string
	Length map[string]int
	Titles map[string]string
	Sequences map[string]string
}

func parseTitle(line []byte) string {
	res := make([]byte, 0)
	count := 0
Loop:
	for i := 1; i < len(line); i++ {
		switch line[i] {
		case 32:
			break Loop
		case 124:
			if count < 1 {
				res = append(res, line[i])
				count++
			} else {
				break Loop
			}
		case 59:
			// ignore the comment line
		default:
			res = append(res, line[i])
		}
	}
	return string(res)
}

// ReadFasta read a local file and parse it
// Only standard format is supported
func ReadFasta(path string) (*Fasta, error){
	inFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	res := Fasta{
		Source: path,
		Length: make(map[string]int),
		Identifiers: make([]string, 0),
		Titles: make(map[string]string),
		Sequences: make(map[string]string),
	}
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) > 0 {
			switch line[0] {
			case 62:
				id := parseTitle(line)
				res.Identifiers = append(res.Identifiers, id)
				res.Titles[id] = string(line)
			default:
				res.Sequences[res.Identifiers[len(res.Identifiers)-1]] += string(line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	for _, id := range res.Identifiers {
		res.Length[id] = len(res.Sequences[id])
	}
	return &res, nil
}