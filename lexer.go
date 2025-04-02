package main

type Lexer struct {
	input        string
	index        int
	readPosition int
	ch           byte
}

func newLexer(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.index = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	begin := l.index

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[begin:l.index]
}

func (l *Lexer) readNumber() string {
	begin := l.index

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[begin:l.index]
}

func (l *Lexer) nextToken() Token {
	var tok Token

	l.skipWhitespace()

	tok.start = l.index

	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			l.readChar()
			tok.tag = equalEqual
		} else {
			tok.tag = equal
		}
	case '+':
		tok.tag = plus
	case '-':
		tok.tag = minus
	case '!':
		if l.peakChar() == '=' {
			l.readChar()
			tok.tag = bangEqual
		} else {
			tok.tag = bang
		}
	case '/':
		tok.tag = slash
	case '*':
		tok.tag = asterisk
	case '<':
		tok.tag = lAngleBracket
	case '>':
		tok.tag = rAngleBracket
	case ',':
		tok.tag = comma
	case ';':
		tok.tag = semicolon
	case '(':
		tok.tag = lParen
	case ')':
		tok.tag = rParen
	case '{':
		tok.tag = lBrace
	case '}':
		tok.tag = rBrace
	case 0:
		tok.tag = eof
	default:
		if isLetter(l.ch) {
			identifier := l.readIdentifier()
			tok.tag = lookupIdent(identifier)
			return tok
		} else if isDigit(l.ch) {
			l.readNumber()
			tok.tag = numberLiteral
			return tok
		} else {
			tok.tag = illegal
		}
	}

	l.readChar()
	return tok
}
