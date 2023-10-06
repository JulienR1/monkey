package lexer

import (
	"github.com/JulienR1/monkey/internal/pkg/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var t token.Token

	lexer.skipWhitespace()

	switch lexer.character {
	case '=':
		if lexer.peekChar() == '=' {
			character := lexer.character
			lexer.readChar()
			t = token.Token{Type: token.EQUAL, Literal: string(character) + string(lexer.character)}
		} else {
			t = newToken(token.ASSIGN, lexer.character)
		}
	case ';':
		t = newToken(token.SEMICOLON, lexer.character)
	case '(':
		t = newToken(token.LPAREN, lexer.character)
	case ')':
		t = newToken(token.RPAREN, lexer.character)
	case '{':
		t = newToken(token.LBRACE, lexer.character)
	case '}':
		t = newToken(token.RBRACE, lexer.character)
	case ',':
		t = newToken(token.COMMA, lexer.character)
	case '+':
		t = newToken(token.PLUS, lexer.character)
	case '-':
		t = newToken(token.MINUS, lexer.character)
	case '!':
		if lexer.peekChar() == '=' {
			character := lexer.character
			lexer.readChar()
			t = token.Token{Type: token.NOT_EQUAL, Literal: string(character) + string(lexer.character)}
		} else {
			t = newToken(token.BANG, lexer.character)
		}
	case '*':
		t = newToken(token.ASTERISK, lexer.character)
	case '/':
		t = newToken(token.SLASH, lexer.character)
	case '<':
		t = newToken(token.LESS_THAN, lexer.character)
	case '>':
		t = newToken(token.GREATER_THAN, lexer.character)

	case 0:
		t = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(lexer.character) {
			literal := lexer.readIdentifier()
			tokenType := token.LookupIdentifier(literal)
			return token.Token{Type: tokenType, Literal: literal}
		} else if isDigit(lexer.character) {
			return token.Token{Type: token.INT, Literal: lexer.readNumber()}
		} else {
			t = token.Token{Type: token.ILLEGAL, Literal: ""}
		}
	}

	lexer.readChar()
	return t
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	}
	return lexer.input[lexer.readPosition]
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.character) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.character) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]

}

func (lexer *Lexer) skipWhitespace() {
	for lexer.character == ' ' || lexer.character == '\t' || lexer.character == '\n' || lexer.character == '\r' {
		lexer.readChar()
	}
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func isLetter(character byte) bool {
	return 'a' <= character && 'z' >= character || 'A' <= character && 'Z' >= character || character == '_'
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}
