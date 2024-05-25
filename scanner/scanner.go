package scanner

import (
	"lox/tokens"
	"unicode"
)

type Scanner struct {
	source               []rune
	start, current, line int
	tokens               []tokens.Token
}

func NewScanner(source string) Scanner {
	return Scanner{
		source:  []rune(source),
		start:   0,
		current: 0,
		line:    1,
		tokens:  make([]tokens.Token, 0, 100),
	}
}

func (s *Scanner) ScanTokens() []tokens.Token {
	for !s.isAtEnd() {
		s.scanToken()
	}
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	switch r := s.advance(); r {
	// Single-character tokens
	case '(':
		s.addToken(tokens.LEFT_PAREN)
	case ')':
		s.addToken(tokens.RIGHT_PAREN)
	case '{':
		s.addToken(tokens.LEFT_BRACE)
	case '}':
		s.addToken(tokens.RIGHT_BRACE)
	case ',':
		s.addToken(tokens.COMMA)
	case '.':
		s.addToken(tokens.DOT)
	case '-':
		s.addToken(tokens.MINUS)
	case '+':
		s.addToken(tokens.PLUS)
	case ';':
		s.addToken(tokens.SEMICOLON)
	case '*':
		s.addToken(tokens.STAR)

	// One or two character tokens. Need to look ahead to
	// distinguish between tokens
	case '!':
		if s.match('=') {
			s.addToken(tokens.BANG_EQUAL)
		} else {
			s.addToken(tokens.BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(tokens.EQUAL_EQUAL)
		} else {
			s.addToken(tokens.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(tokens.LESS_EQUAL)
		} else {
			s.addToken(tokens.LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(tokens.GREATER_EQUAL)
		} else {
			s.addToken(tokens.GREATER)
		}

	case '/':
		if s.match('/') {
			for !s.isAtEnd() && s.peek() != '\n' {
				s.advance()
			}
		} else {
			s.addToken(tokens.SLASH)
		}

	// Ignore whitespace and new lines
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++

	case '"':
		s.addStringToken()

	default:
		// number literal
		if unicode.IsDigit(r) {
			s.addNumberToken()
		}
		error(s.line, "Unexpected character.")
	}
}

func (s *Scanner) addToken(tokenType tokens.TokenType) {
	token := tokens.Token{
		Type:   tokenType,
		Lexeme: string(s.source[s.start:s.current]),
		Line:   s.line,
	}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) advance() rune {
	r := s.source[s.current]
	s.current++
	return r
}

// Conditional advance. Only advance if the current rune
// matches the expected rune
func (s *Scanner) match(expected rune) bool {
	if s.peek() == expected {
		s.advance()
		return true
	}
	return false
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\x00'
	} else {
		return s.source[s.current]
	}
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\x00'
	} else {
		return s.source[s.current+1]
	}
}

func (s *Scanner) addStringToken() {
	for !s.isAtEnd() && s.peek() != '"' {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		error(s.line, "Unterminated string.")
	}

	token := tokens.Token{
		Type:   tokens.STRING,
		Lexeme: string(s.source[s.start+1 : s.current]),
		Line:   s.line,
	}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) addNumberToken() {
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && unicode.IsDigit(s.peekNext()) {
		s.advance()
	}

	for unicode.IsDigit(s.peek()) {
		s.advance()
	}

	s.addToken(tokens.NUMBER)
}
