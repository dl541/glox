package scanner

import (
	"lox/tokens"
	"reflect"
	"testing"
)

func TestSingleToken(t *testing.T) {
	tests := []struct {
		name          string
		source        string
		expectedToken tokens.Token
	}{
		{"scan single token", "{", tokens.Token{
			Type:   tokens.LEFT_BRACE,
			Lexeme: "{",
			Line:   1,
		}},
		{"scan bang", "!", tokens.Token{
			Type:   tokens.BANG,
			Lexeme: "!",
			Line:   1,
		}},
		{"scan bang equal", "!=", tokens.Token{
			Type:   tokens.BANG_EQUAL,
			Lexeme: "!=",
			Line:   1,
		}},
		{"scan number", "123.0", tokens.Token{
			Type:   tokens.NUMBER,
			Lexeme: "123.0",
			Line:   1,
		}},
		{"scan string", "\"abc\"", tokens.Token{
			Type:   tokens.STRING,
			Lexeme: "abc",
			Line:   1,
		}},
		{"scan identifier", "and123", tokens.Token{
			Type:   tokens.IDENTIFIER,
			Lexeme: "and123",
			Line:   1,
		}},
		{"scan reserved keyword", "return", tokens.Token{
			Type:   tokens.RETURN,
			Lexeme: "return",
			Line:   1,
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scanner := NewScanner(test.source)
			res := scanner.ScanTokens()
			if !reflect.DeepEqual(res[0], test.expectedToken) {
				t.Errorf("Expect token %+v, got %+v", res[0], test.expectedToken)
			}

			if res[1].Type != tokens.EOF {
				t.Errorf("Expect EOF token, got %+v", res[len(res)-1])
			}

			if !scanner.isAtEnd() {
				t.Errorf("Expect scanner to finish scanning. Stuck at index %v", scanner.current)
			}
		})
	}
}

func TestScanBlankLine(t *testing.T) {
	scanner := NewScanner("\n")
	res := scanner.ScanTokens()
	expectedToken := tokens.Token{
		Type:   tokens.EOF,
		Lexeme: "",
		Line:   2,
	}
	if !reflect.DeepEqual(res[0], expectedToken) {
		t.Errorf("Expected EOF token %#v, got %#v", expectedToken, res[0])
	}
	if !scanner.isAtEnd() {
		t.Errorf("Expect scanner to finish scanning. Stuck at index %v", scanner.current)
	}
}

func TestScanMultipleTokens(t *testing.T) {
	tests := []struct {
		name           string
		source         string
		expectedTokens []tokens.Token
	}{
		{"scan numerical expression", "123.0 + 23\n-235/3", []tokens.Token{
			{Type: tokens.NUMBER, Lexeme: "123.0", Line: 1},
			{Type: tokens.PLUS, Lexeme: "+", Line: 1},
			{Type: tokens.NUMBER, Lexeme: "23", Line: 1},
			{Type: tokens.MINUS, Lexeme: "-", Line: 2},
			{Type: tokens.NUMBER, Lexeme: "235", Line: 2},
			{Type: tokens.SLASH, Lexeme: "/", Line: 2},
			{Type: tokens.NUMBER, Lexeme: "3", Line: 2},
			{Type: tokens.EOF, Lexeme: "", Line: 2},
		}},
		{"reserved keywords with identifiers", "a and b or \"AND\" or FUN", []tokens.Token{
			{Type: tokens.IDENTIFIER, Lexeme: "a", Line: 1},
			{Type: tokens.AND, Lexeme: "and", Line: 1},
			{Type: tokens.IDENTIFIER, Lexeme: "b", Line: 1},
			{Type: tokens.OR, Lexeme: "or", Line: 1},
			{Type: tokens.STRING, Lexeme: "AND", Line: 1},
			{Type: tokens.OR, Lexeme: "or", Line: 1},
			{Type: tokens.IDENTIFIER, Lexeme: "FUN", Line: 1},
			{Type: tokens.EOF, Lexeme: "", Line: 1},
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scanner := NewScanner(test.source)
			res := scanner.ScanTokens()
			if !reflect.DeepEqual(res, test.expectedTokens) {
				t.Errorf("Expected tokens %+v, got %+v", test.expectedTokens, res)
			}
			if !scanner.isAtEnd() {
				t.Errorf("Expect scanner to finish scanning. Stuck at index %v", scanner.current)
			}
		})
	}
}

func TestScanComments(t *testing.T) {
	scanner := NewScanner("a/b// comments should be skipped \n 1.3/123")
	res := scanner.ScanTokens()
	expectedTokens := []tokens.Token{
		{Type: tokens.IDENTIFIER, Lexeme: "a", Line: 1},
		{Type: tokens.SLASH, Lexeme: "/", Line: 1},
		{Type: tokens.IDENTIFIER, Lexeme: "b", Line: 1},
		{Type: tokens.NUMBER, Lexeme: "1.3", Line: 2},
		{Type: tokens.SLASH, Lexeme: "/", Line: 2},
		{Type: tokens.NUMBER, Lexeme: "123", Line: 2},
		{Type: tokens.EOF, Lexeme: "", Line: 2},
	}

	if !reflect.DeepEqual(res, expectedTokens) {
		t.Errorf("Expected tokens %+v, got %+v", expectedTokens, res)
	}
	if !scanner.isAtEnd() {
		t.Errorf("Expect scanner to finish scanning. Stuck at index %v", scanner.current)
	}
}
