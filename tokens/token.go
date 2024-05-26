package tokens

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
}
