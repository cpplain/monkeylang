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
		expectedTag       TokenTag
		expectedFirstChar string
	}{
		{keywordLet, "l"},
		{identifier, "f"},
		{equal, "="},
		{numberLiteral, "5"},
		{semicolon, ";"},
		{keywordLet, "l"},
		{identifier, "t"},
		{equal, "="},
		{numberLiteral, "1"},
		{semicolon, ";"},
		{keywordLet, "l"},
		{identifier, "a"},
		{equal, "="},
		{keywordFn, "f"},
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
		{keywordLet, "l"},
		{identifier, "r"},
		{equal, "="},
		{identifier, "a"},
		{lParen, "("},
		{identifier, "f"},
		{comma, ","},
		{identifier, "t"},
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
		{numberLiteral, "1"},
		{rAngleBracket, ">"},
		{numberLiteral, "5"},
		{semicolon, ";"},
		{keywordIf, "i"},
		{lParen, "("},
		{numberLiteral, "5"},
		{lAngleBracket, "<"},
		{numberLiteral, "1"},
		{rParen, ")"},
		{lBrace, "{"},
		{keywordReturn, "r"},
		{keywordTrue, "t"},
		{semicolon, ";"},
		{rBrace, "}"},
		{keywordElse, "e"},
		{lBrace, "{"},
		{keywordReturn, "r"},
		{keywordFalse, "f"},
		{semicolon, ";"},
		{rBrace, "}"},
		{numberLiteral, "1"},
		{equalEqual, "="},
		{numberLiteral, "1"},
		{semicolon, ";"},
		{numberLiteral, "1"},
		{bangEqual, "!"},
		{numberLiteral, "9"},
		{semicolon, ";"},
		{eof, ""},
	}

	l := newLexer(input)

	for i, tt := range test {
		tok := l.nextToken()

		if tok.tag != tt.expectedTag {
			t.Fatalf("test[%d]: wrong tag: expected='%s', actual='%s'",
				i, tagString[tt.expectedTag], tagString[tok.tag])
		}

		if tok.tag != eof && string(l.input[tok.start]) != tt.expectedFirstChar {
			t.Fatalf("test[%d]: wrong first char: expected='%s', actual='%s'",
				i, tt.expectedFirstChar, string(l.input[tok.start]))
		}
	}
}
