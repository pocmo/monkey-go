package lexer

import (
	"fmt"
	"testing"

	"github.com/pocmo/monkey-go/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, expectedToken := range tests {
		token := lexer.NextToken()

		fmt.Printf("Test[%d]: %q\n", i, lexer.ch)

		if token.Type != expectedToken.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, expectedToken.expectedType, token.Type)
		}

		if token.Literal != expectedToken.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q", i, expectedToken.expectedLiteral, token.Literal)
		}
	}
}
