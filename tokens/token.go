package tokens

import "fmt"

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
}

func (token Token) String() string {
	return fmt.Sprintf("%v %v %v", token.Type, token.Lexeme, token.Line)
}
