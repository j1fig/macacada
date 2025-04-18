package lexer

import (
	"fmt"
	"testing"

	"monkey/token"
)


func logLexer(l *Lexer) {
	fmt.Printf("pos: %d\nreadpos: %d\nch: %c\n", l.position, l.readPosition, l.ch)
}

func logToken(t token.Token) {
	fmt.Printf("type: %s\nliteral: %s\n", t.Type, t.Literal)
}


func TestNewToken(t *testing.T) {
	input := "=+(){},;"

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

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			logToken(tok)
			logLexer(l)
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q. got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			logToken(tok)
			logLexer(l)
			t.Fatalf("tests[%d] - literal wrong. expected=%q. got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
