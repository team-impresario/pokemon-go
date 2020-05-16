package scanner

import (
	"bufio"
	"bytes"
	"errors"
	// "fmt"
	"io"
)

type Token int

var keywordMap = map[string]string{
	"exeggcutor": "func",
	"goldeen":    "import",
	"klefki":     "case",
	"alakazam":   "goto",
	"snorlax":    "break",
	"clink":      "continue",
	"geodude":    "const",
	"jynx":       "var",
	"squirtle":   "default",
	"psyduck":    "chan",
	"sandshrew":  "fallthrough",
	"forretress": "for",
	"chansey":    "if",
	"slowbro":    "package",
	"unown":      "else",
}

var eof = rune(0)

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

func (s *Scanner) Scan() (lit string, err error) {
	ch := s.read()
	if isLetter(ch) {
		s.unread()
		return s.scanKeyword(), nil
	} else if isQuote(ch) {
		s.unread()
		return s.scanAll(), nil
	} else if ch == eof {
		return "", nil
	} else {
		s.unread()
		return s.scanRest(), nil
	}
	return "", errors.New("Your code doesn't work for reasons of incomprehensibility. Have you considered quitting programming?")
}

func (s *Scanner) scanRest() (lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	return buf.String()
}

func (s *Scanner) scanAll() (lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isQuote(ch) {
			buf.WriteRune(ch)
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func (s *Scanner) scanKeyword() (lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	if val, ok := keywordMap[buf.String()]; ok {
		return val
	} else {
		return buf.String()
	}
}
