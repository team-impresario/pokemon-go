package parser

import (
	"bytes"
	"io"
	"log"

	scanner "pokemon-go/parser/scanner"
)

type Parser struct {
	s      *scanner.Scanner
	parsed bytes.Buffer
}

func (p *Parser) scan() (s string) {

	scanned, err := p.s.Scan()
	if err != nil {
		log.Fatal(err)
	}
	_, _ = p.parsed.WriteString(scanned)

	return scanned
}

func (p *Parser) Parse() string {
	for {
		scanned := p.scan()
		if scanned == "" {
			break
		}
	}
	return p.parsed.String()
}

func NewParser(r io.Reader) *Parser {
	var buf bytes.Buffer
	return &Parser{
		s:      scanner.NewScanner(r),
		parsed: buf,
	}
}
