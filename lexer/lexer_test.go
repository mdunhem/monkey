package lexer

import (
	"testing"

	"github.com/mdunhem/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LEFT_PAREN, "("},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_CURLY, "{"},
		{token.RIGHT_CURLY, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := NewLexer(input)

	for i, test := range tests {
		currentToken := lexer.NextToken()

		if currentToken.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong, expected: '%q' -- got: '%q'", i, test.expectedType, currentToken.Type)
		}

		if currentToken.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected: '%q' -- got: '%q'", i, test.expectedLiteral, currentToken.Literal)
		}
	}
}
