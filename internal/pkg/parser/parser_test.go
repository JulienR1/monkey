package parser

import (
	"testing"

	"github.com/JulienR1/monkey/internal/pkg/ast"
	"github.com/JulienR1/monkey/internal/pkg/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
    let x = 5;
    let y = 10;
    let foobar = 63366;
    `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	expectedIdentifiers := []string{"x", "y", "foobar"}
	for index, expectedIdentifer := range expectedIdentifiers {
		statement := program.Statements[index]
		if !testLetStatement(t, statement, expectedIdentifer) {
			return
		}
	}

}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, statement ast.Statement, expectedIdentifier string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiteral() not 'let'. got=%q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement not '*ast.LetStatement'. got=%T", statement)
		return false
	}

	if letStatement.Name.Value != expectedIdentifier {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", expectedIdentifier, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != expectedIdentifier {
		t.Errorf("letStatement.TokenLiteral() not '%s'. got=%s", expectedIdentifier, letStatement.TokenLiteral())
		return false
	}

	return true
}
