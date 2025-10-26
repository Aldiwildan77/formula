package engine

import (
	"fmt"
	"strings"
	"unicode"
)

// We do a tokenizer in here
type TokenType int

const (
	TokenIdent TokenType = iota
	TokenNumber
	TokenString
	TokenLParen
	TokenRParen
	TokenComma
	TokenTrue
	TokenFalse
	TokenEOF
	TokenComment
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input    []rune
	pos      int
	debugger *Debugger
}

func NewLexer(input string, debugger *Debugger) *Lexer {
	return &Lexer{
		input:    []rune(strings.TrimSpace(input)),
		debugger: debugger,
	}
}

func (l *Lexer) skipSpaces() {
	for l.pos < len(l.input) && unicode.IsSpace(l.input[l.pos]) {
		l.pos++
	}
}

func (l *Lexer) NextToken() Token {
	l.skipSpaces()

	if l.pos >= len(l.input) {
		return Token{
			Type: TokenEOF,
		}
	}

	character := l.input[l.pos]

	switch {
	case character == '#':
		// Handle single-line comment with '#'
		start := l.pos
		l.pos++
		for l.pos < len(l.input) && l.input[l.pos] != '\n' {
			l.pos++
		}
		comment := string(l.input[start:l.pos])
		token := Token{Type: TokenComment, Value: comment}
		if l.debugger != nil {
			l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
		}
		return token
	case character == '-' && l.pos+1 < len(l.input) && l.input[l.pos+1] == '-':
		// Handle single-line comment with '--'
		start := l.pos
		l.pos += 2
		for l.pos < len(l.input) && l.input[l.pos] != '\n' {
			l.pos++
		}
		comment := string(l.input[start:l.pos])
		token := Token{Type: TokenComment, Value: comment}
		if l.debugger != nil {
			l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
		}
		return token
	case character == '/' && l.pos+1 < len(l.input) && l.input[l.pos+1] == '*':
		// Handle multi-line comment /* ... */
		start := l.pos
		l.pos += 2
		for l.pos+1 < len(l.input) && !(l.input[l.pos] == '*' && l.input[l.pos+1] == '/') {
			l.pos++
		}
		if l.pos+1 < len(l.input) {
			l.pos += 2 // Skip the closing */
		}
		comment := string(l.input[start:l.pos])
		token := Token{Type: TokenComment, Value: comment}
		if l.debugger != nil {
			l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
		}
		return token
	case unicode.IsLetter(character):
		start := l.pos
		for l.pos < len(l.input) && (unicode.IsLetter(l.input[l.pos]) || l.input[l.pos] == '_') {
			l.pos++
		}
		word := strings.ToUpper(string(l.input[start:l.pos]))
		switch word {
		case "TRUE":
			token := Token{Type: TokenTrue, Value: "TRUE"}
			if l.debugger != nil {
				if l.debugger != nil {
					l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
				}
			}
			return token
		case "FALSE":
			token := Token{Type: TokenFalse, Value: "False"}
			if l.debugger != nil {
				l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
			}
			return token
		default:
			token := Token{Type: TokenIdent, Value: word}
			if l.debugger != nil {
				l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
			}
			return token
		}
	case unicode.IsDigit(character):
		start := l.pos
		for l.pos < len(l.input) && (unicode.IsDigit(l.input[l.pos])) {
			l.pos++
		}
		token := Token{Type: TokenNumber, Value: string(l.input[start:l.pos])}
		if l.debugger != nil {
			l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
		}
		return token
	case character == '\'' || character == '"':
		quote := character
		l.pos++
		start := l.pos
		for l.pos < len(l.input) && l.input[l.pos] != quote {
			l.pos++
		}
		str := string(l.input[start:l.pos])
		l.pos++
		token := Token{Type: TokenString, Value: str}
		if l.debugger != nil {
			l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
		}
		return token
	case character == '(':
		l.pos++
		token := Token{Type: TokenLParen, Value: "("}
		if l.debugger != nil {
			l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
		}
		return token
	case character == ')':
		l.pos++
		token := Token{Type: TokenRParen, Value: ")"}
		if l.debugger != nil {
			l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
		}
		return token
	case character == ',':
		l.pos++
		token := Token{Type: TokenComma, Value: ","}
		if l.debugger != nil {
			l.debugger.Log(PhaseLexer, 0, fmt.Sprintf("Token: %v", token))
		}
		return token
	default:
		l.pos++
		return l.NextToken()
	}
}
