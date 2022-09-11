package token

type TokenType string

const (
	ILLEGAL = "Illegal"
	EOF = "EOF"
	// Identifier + Literals
	IDENT = "IDENT"
	INT = "INT"
	// Operators
	ASSIGN = "="
	PLUS = "+"
	// Delimiter
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

