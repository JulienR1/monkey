package parser

import (
	"fmt"

	"github.com/JulienR1/monkey/internal/pkg/ast"
	"github.com/JulienR1/monkey/internal/pkg/lexer"
	"github.com/JulienR1/monkey/internal/pkg/token"
)

type Parser struct {
	lexer *lexer.Lexer

	errors []string

	currentToken token.Token
	peekToken    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer, errors: []string{}}

	parser.NextToken()
	parser.NextToken()

	return parser
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(tokenType token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tokenType, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) NextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}

	for !parser.currentTokenIs(token.EOF) {
		statement := parser.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.NextToken()
	}

	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}

// Rule: let <identifier> = <expression>;
func (parser *Parser) parseLetStatement() *ast.LetStatement {
	statement := ast.LetStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.IDENTIFIER) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skip until semicolon (for now)
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.NextToken()
	}

	return &statement
}

func (parser *Parser) currentTokenIs(tokenType token.TokenType) bool {
	return parser.currentToken.Type == tokenType
}

func (parser *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return parser.peekToken.Type == tokenType
}

func (parser *Parser) expectPeek(tokenType token.TokenType) bool {
	if parser.peekTokenIs(tokenType) {
		parser.NextToken()
		return true
	}
	parser.peekError(tokenType)
	return false
}
