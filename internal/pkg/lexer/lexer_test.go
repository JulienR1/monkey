package lexer

import (
	"testing"

	"github.com/JulienR1/monkey/internal/pkg/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
	}

	testInput(t, input, tests)
}

func TestNextToken2(t *testing.T) {
	input := `
let five = 5;
let ten = 10;

let add = fn(x,y) { x + y; }

let result = add(five, ten);
`

	tests := []token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "five"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "ten"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "add"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.FUNCTION, Literal: "fn"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "result"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENTIFIER, Literal: "add"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "five"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "ten"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},
	}

	testInput(t, input, tests)
}

func TestNextToken3(t *testing.T) {
	input := `
   !-/*5;
   5 < 10 > 5; 
    `

	tests := []token.Token{
		{Type: token.BANG, Literal: "!"},
		{Type: token.MINUS, Literal: "-"},
		{Type: token.SLASH, Literal: "/"},
		{Type: token.ASTERISK, Literal: "*"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.INT, Literal: "5"},
		{Type: token.LESS_THAN, Literal: "<"},
		{Type: token.INT, Literal: "10"},
		{Type: token.GREATER_THAN, Literal: ">"},
		{Type: token.INT, Literal: "5"},
	}

	testInput(t, input, tests)
}

func TestNextToken4(t *testing.T) {
	input := `
   if ( 5 < 10 ) {
        return true;
   } else {
        return false;
   }
    `

	tests := []token.Token{
		{Type: token.IF, Literal: "if"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.INT, Literal: "5"},
		{Type: token.LESS_THAN, Literal: "<"},
		{Type: token.INT, Literal: "10"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.ELSE, Literal: "else"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.FALSE, Literal: "false"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
	}

	testInput(t, input, tests)
}

func TestNextToken5(t *testing.T) {
	input := `
        10 == 9
        10 != 10
        `

	tests := []token.Token{
		{Type: token.INT, Literal: "10"},
		{Type: token.EQUAL, Literal: "=="},
		{Type: token.INT, Literal: "9"},
		{Type: token.INT, Literal: "10"},
		{Type: token.NOT_EQUAL, Literal: "!="},
		{Type: token.INT, Literal: "10"},
	}

	testInput(t, input, tests)
}

func testInput(t *testing.T, input string, expectedTokens []token.Token) {
	lexer := New(input)

	for index, expected := range expectedTokens {
		token := lexer.NextToken()

		if token.Type != expected.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected %q, got %q", index, expected.Type, token.Type)
		}

		if token.Literal != expected.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected %q, got %q", index, expected.Literal, token.Literal)
		}
	}
}
