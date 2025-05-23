package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	ASTERISK   = "*"
	SLASH    = "/"

	// Boolean operators
	BANG   = "!"
	EQ     = "=="
	NOT_EQ = "!="
	GT	   = ">"
	LT     = "<"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE	 = "TRUE"
	FALSE	 = "FALSE"
	IF       = "IF"
	ELSE	 = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(keyword string) TokenType {
	if tokenType, ok := keywords[keyword]; ok {
		return tokenType
	}
	return IDENT
}
