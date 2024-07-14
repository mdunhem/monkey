package lexer

import "github.com/mdunhem/monkey/token"

type Lexer struct {
	input           string
	currentPosition int
	readPosition    int
	char            byte
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readNextChar()

	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var _token token.Token

	switch lexer.char {
	case '=':
		_token = token.Token{Type: token.ASSIGN, Literal: string(lexer.char)}
	case '+':
		_token = token.Token{Type: token.PLUS, Literal: string(lexer.char)}
	case ',':
		_token = token.Token{Type: token.COMMA, Literal: string(lexer.char)}
	case ';':
		_token = token.Token{Type: token.SEMICOLON, Literal: string(lexer.char)}
	case '(':
		_token = token.Token{Type: token.LEFT_PAREN, Literal: string(lexer.char)}
	case ')':
		_token = token.Token{Type: token.RIGHT_PAREN, Literal: string(lexer.char)}
	case '{':
		_token = token.Token{Type: token.LEFT_CURLY, Literal: string(lexer.char)}
	case '}':
		_token = token.Token{Type: token.RIGHT_CURLY, Literal: string(lexer.char)}
	case 0:
		_token = token.Token{Type: token.EOF, Literal: ""}
	}

	lexer.readNextChar()

	return _token
}

func (lexer *Lexer) readNextChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}

	lexer.currentPosition = lexer.readPosition
	lexer.readPosition += 1
}
