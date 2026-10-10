package cli

import (
	"bytes"
	"fmt"
)

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenEOL
	TokenSkip
	TokenText
	TokenOpen
	TokenClose

	TokenIf
	TokenYield
	TokenTry
	TokenLoop
	TokenWhile
)

type Token struct {
	Type     TokenType
	Value    string
	Location Location
}

type Location struct {
	Index  int
	Line   int
	Length int
}

func (l Location) IsEmpty() bool {
	return l.Index == 0 && l.Length == 0
}

func (l Location) String() string {
	return fmt.Sprintf("%d; %d", l.Line+1, l.Index+1)
}

type TokenReader struct {
	Source string
	Index  int
	Line   int
}

func (r *TokenReader) ReadToEnd() (res []Token) {
readToEndLoop:
	for {
		tok := r.Next()

		switch tok.Type {
		case TokenSkip:
			continue
		case TokenEOF:
			break readToEndLoop
		}

		res = append(res, tok)
	}
	return
}

func (r *TokenReader) Next() Token {
	loc := Location{
		Index: r.Index,
		Line:  r.Line,
	}
	t := r.next()
	loc.Length = r.Index - loc.Index
	t.Location = loc
	return t
}

func (r *TokenReader) next() Token {
	chr := r.peekChar()

	r.moveIndex()

	quote := byte(0)

	switch chr {
	case 0:
		return Token{Type: TokenEOF}
	case '[':
		return Token{Type: TokenOpen}
	case ']':
		return Token{Type: TokenClose}
	case ';', '\n':
		return Token{Type: TokenEOL}
	case ' ', '\r', '\t':
		return Token{Type: TokenSkip}
	case '"', '\'':
		quote = chr
	case '\\':
		if r.peekChar() == '\n' {

			r.moveIndex()
			return Token{Type: TokenSkip}
		}
	}

	if quote == 0 {
		r.Index--
	}

	var buff bytes.Buffer

textLoop:
	for {
		esc := false
		chr = r.peekChar()
		r.moveIndex()
		switch chr {
		case '\r':
			continue
		case '\n', 0:
			r.Index--
			break textLoop
		case ' ', ']', ';':
			if quote == 0 {
				r.Index--
				break textLoop
			}
		case '\\':
			esc = true
			chr = r.peekChar()

			r.moveIndex()

			switch chr {
			case '\n':
				continue
			case 'n':
				chr = '\n'
			case 'r':
				chr = '\r'
			case 't':
				chr = '\t'
			}
		}
		if chr == quote && !esc {
			break
		}
		buff.WriteByte(chr)
	}

	text := buff.String()

	if quote == 0 {
		switch text {
		case "If":
			return Token{Type: TokenIf}
		case "Yield":
			return Token{Type: TokenYield}
		case "Try":
			return Token{Type: TokenTry}
		case "Loop":
			return Token{Type: TokenLoop}
		case "While":
			return Token{Type: TokenWhile}
		}
	}

	return Token{Type: TokenText, Value: text}
}

func (r *TokenReader) peekChar() byte {
	if r.Index >= len(r.Source) {
		return 0
	}
	return r.Source[r.Index]
}

func (r *TokenReader) moveIndex() {
	if r.peekChar() == '\n' {
		r.Line++
	}
	r.Index++
}
