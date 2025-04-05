package main

type TokenTag int

type Token struct {
	tag   TokenTag
	start int
	end   int
}

const (
	illegal TokenTag = iota
	eof
	identifier
	numberLiteral
	equal
	plus
	minus
	bang
	asterisk
	slash
	lAngleBracket
	rAngleBracket
	equalEqual
	bangEqual
	comma
	semicolon
	lParen
	rParen
	lBrace
	rBrace
	keywordFn
	keywordLet
	keywordTrue
	keywordFalse
	keywordIf
	keywordElse
	keywordReturn
)

var tagString = map[TokenTag]string{
	illegal:       "illegal",
	eof:           "EOF",
	identifier:    "identifier",
	numberLiteral: "numberLiteral",
	equal:         "equal",
	plus:          "plus",
	minus:         "minus",
	bang:          "bang",
	asterisk:      "asterisk",
	slash:         "slash",
	lAngleBracket: "lAngleBracket",
	rAngleBracket: "rAngleBracket",
	equalEqual:    "equalEqual",
	bangEqual:     "bangEqual",
	comma:         "comma",
	semicolon:     "semicolon",
	lParen:        "lParen",
	rParen:        "rParen",
	lBrace:        "lBrace",
	rBrace:        "rBrace",
	keywordFn:     "keywordFn",
	keywordLet:    "keywordLet",
	keywordTrue:   "keywordTrue",
	keywordFalse:  "keywordFalse",
	keywordIf:     "keywordIf",
	keywordElse:   "keywordElse",
	keywordReturn: "keywordReturn",
}

var keywords = map[string]TokenTag{
	"fn":     keywordFn,
	"let":    keywordLet,
	"true":   keywordTrue,
	"false":  keywordFalse,
	"if":     keywordIf,
	"else":   keywordElse,
	"return": keywordReturn,
}

func lookupIdent(ident string) TokenTag {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return identifier
}
