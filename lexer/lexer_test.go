package lexer

import (
	"testing"

	"ta/token"
)

func TestNextToken(t *testing.T) {
	input := `!=;
	==;
	let x = 5;
	
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NOT_EQ, "!="},
		{token.SEMICOLON, ";"},
		{token.EQ, "=="},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
