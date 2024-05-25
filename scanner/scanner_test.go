package scanner

import (
	"lox/tokens"
	"testing"
)

func TestScanTokens(t *testing.T) {
	scanner := NewScanner("123.0")
	res := scanner.ScanTokens()
	token := res[0]
	expectedToken := tokens.Token{
		Type:   tokens.NUMBER,
		Lexeme: "123.0",
		Line:   1,
	}
	if token != expectedToken {
		t.Errorf("Expect token %#v, got %#v", token, expectedToken)
	}

	if !scanner.isAtEnd() {
		t.Errorf("Expect scanner to finish scanning. Stuck at index %v", scanner.current)
	}
}
