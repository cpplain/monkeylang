package main

import "testing"

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
    x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
    return true;
} else {
    return false;
}

10 == 10;
10 != 9;
`

	test := []struct {
		expectedTag    TokenTag
		expectedString string
	}{
		{keywordLet, "let"},
		{identifier, "five"},
		{equal, "="},
		{numberLiteral, "5"},
		{semicolon, ";"},
		{keywordLet, "let"},
		{identifier, "ten"},
		{equal, "="},
		{numberLiteral, "10"},
		{semicolon, ";"},
		{keywordLet, "let"},
		{identifier, "add"},
		{equal, "="},
		{keywordFn, "fn"},
		{lParen, "("},
		{identifier, "x"},
		{comma, ","},
		{identifier, "y"},
		{rParen, ")"},
		{lBrace, "{"},
		{identifier, "x"},
		{plus, "+"},
		{identifier, "y"},
		{semicolon, ";"},
		{rBrace, "}"},
		{semicolon, ";"},
		{keywordLet, "let"},
		{identifier, "result"},
		{equal, "="},
		{identifier, "add"},
		{lParen, "("},
		{identifier, "five"},
		{comma, ","},
		{identifier, "ten"},
		{rParen, ")"},
		{semicolon, ";"},
		{bang, "!"},
		{minus, "-"},
		{slash, "/"},
		{asterisk, "*"},
		{numberLiteral, "5"},
		{semicolon, ";"},
		{numberLiteral, "5"},
		{lAngleBracket, "<"},
		{numberLiteral, "10"},
		{rAngleBracket, ">"},
		{numberLiteral, "5"},
		{semicolon, ";"},
		{keywordIf, "if"},
		{lParen, "("},
		{numberLiteral, "5"},
		{lAngleBracket, "<"},
		{numberLiteral, "10"},
		{rParen, ")"},
		{lBrace, "{"},
		{keywordReturn, "return"},
		{keywordTrue, "true"},
		{semicolon, ";"},
		{rBrace, "}"},
		{keywordElse, "else"},
		{lBrace, "{"},
		{keywordReturn, "return"},
		{keywordFalse, "false"},
		{semicolon, ";"},
		{rBrace, "}"},
		{numberLiteral, "10"},
		{equalEqual, "=="},
		{numberLiteral, "10"},
		{semicolon, ";"},
		{numberLiteral, "10"},
		{bangEqual, "!="},
		{numberLiteral, "9"},
		{semicolon, ";"},
		{eof, ""},
	}

	l := newLexer(input)
	l.tokenize()

	for i, tt := range test {
		tok := l.tokens[i]
		if tok.tag != tt.expectedTag {
			t.Fatalf("test[%d]: wrong tag: expected='%s', actual='%s'",
				i, tagString[tt.expectedTag], tagString[tok.tag])
		}

		if tok.tag != eof && l.input[tok.start:tok.end+1] != tt.expectedString {
			t.Fatalf("test[%d]: wrong string: expected='%s', actual='%s'",
				i, tt.expectedString, l.input[tok.start:tok.end+1])
		}
	}
}
