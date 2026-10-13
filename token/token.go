package token

type TokenType string

type Token struct {
	Type TokenType
	Litteral string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifiers, literals…
	IDENT = "IDENT" // add, foobar, x, y…
	INT = "INT"

	// operators
	ASSIGN = "="
	PLUS = "+"

	// delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTIION"
	LET = "LET"
)