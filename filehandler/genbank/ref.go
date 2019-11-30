package genbank

import (
	"fmt"
	"strings"
)

// Reference about the genome
type Reference struct {
	Description string   `json:"description,omitempty"`
	Authors     []string `json:"authors,omitempty"`
	Title       string   `json:"title,omitempty"`
	Journal     string   `json:"journal,omitempty"`
	PubMed      string   `json:"pubmed,omitempty"`
}

func (ref Reference) String() string {
	res := strings.Builder{}
	res.WriteString(fmt.Sprintln("Description: ", ref.Description))
	res.WriteString(fmt.Sprintln("Authors: ", strings.Join(ref.Authors, " ")))
	res.WriteString(fmt.Sprintln("Title: ", ref.Title))
	res.WriteString(fmt.Sprintln("Journal: ", ref.Journal))
	res.WriteString(fmt.Sprintln("PubMed ID: ", ref.PubMed))
	return res.String()
}

func newReference(a []string) Reference {
	ref := &Reference{}
	ref.Description = a[0]
	var line string
	cur := holder{
		data: make([]string, 0, 1024),
	}
	for i, n := 1, len(a); i < n; i++ {
		line = a[i]
		head, body := strings.Trim(line[:12], " "), line[12:]
		if head == "" {
			cur.data = append(cur.data, body)
			continue
		}
		if cur.name != "" {
			ref.setAttribute(cur)
		}

		cur.name = head
		cur.data = cur.data[:1]
		cur.data[0] = body
	}

	if cur.name != "" {
		ref.setAttribute(cur)
	}
	return *ref
}

func (ref *Reference) setAttribute(cur holder) {
	switch cur.name {
	case "AUTHORS":
		authors := make([]string, 0)
		for _, d := range cur.data {
			authors = append(authors, strings.Split(d, ", ")...)
		}
		m := len(authors)
		authors = append(authors[:m-1], strings.Split(authors[m-1], " and ")...)
		ref.Authors = authors
	case "TITLE":
		ref.Title = strings.Join(cur.data, " ")
	case "JOURNAL":
		ref.Journal = strings.Join(cur.data, " ")
	case "PUBMED":
		ref.PubMed = cur.data[0]
	default:
		println("bypass this line: ", strings.Join(cur.data, " "))
	}
}
