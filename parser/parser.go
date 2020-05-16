package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Token int

const (
	// Keywords
	EXEGGCUTE = "EXEGGCUTE"
	EXEGGUTOR = "EXEGGUTOR"
	KLEFKI    = "KLEFKI"
	ALAKAZAM  = "ALAKAZAM"
	SNORLAX   = "SNORLAX"
	CLINK
	GOLEM
	JYNX
	SQUIRTLE
	PSYDUCK
	KAKUNA
	SANDSHREW
	FORRETRESS
	CHANSEY
	UNOWN
)

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isQuote(ch rune) bool {
	return ch == '"'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhiteSpace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if isQuote(ch) {
		s.unread()
		return s.scanString()
	} else if ch == eof {
		return EOF, ""
	}
	return ch, string
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

func (s *Scanner) scanString() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isQuote(ch) {
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch strings.ToUpper(buf.String()) {
	case "UNOWN":
		return UNOWN, buf.String()
	}
	return IDENT, buf.String()
}
