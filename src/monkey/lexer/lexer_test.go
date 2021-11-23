package lexer

import (
	"fmt"
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="}, {token.PLUS, "+"}, {token.LPAREN, "("}, {token.RPAREN, ")"}, {token.LBRACE, "{"}, {token.RBRACE, "}"}, {token.COMMA, ","}, {token.SEMICOLON, ";"}, {token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.nextToken()
		fmt.Printf("token = %v\n", tok.Type)
		fmt.Printf("tt = %v\n", tt)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
