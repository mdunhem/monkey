package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL     = "ILLEGAL"
	EOF         = "EOF"
	IDENTIFIER  = "IDENTIFIER"
	INT         = "INT"
	ASSIGN      = "="
	PLUS        = "+"
	COMMA       = ","
	SEMICOLON   = ";"
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_CURLY  = "{"
	RIGHT_CURLY = "}"
	FUNCTION    = "FUNCTION"
	LET         = "LET"
)
