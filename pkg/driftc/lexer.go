package driftc

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Lexer struct {
	ParseComments  bool
	ParseAllErrors bool

	input []rune

	pos    int
	line   int
	column int
}

type LexerError struct {
	Token *Token
	Err   error
}

func (e *LexerError) Error() string {
	return fmt.Sprintf("%d:%d: %s", e.Token.Line, e.Token.Column, e.Err)
}

func (l *Lexer) advance() {
	if l.pos >= len(l.input) {
		return
	}
	ch := l.input[l.pos]

	if ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}

	l.pos++
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) readWhile(f func(rune) bool) string {
	var builder strings.Builder

	for {
		ch := l.peek()
		if ch == 0 || !f(ch) {
			return builder.String()
		}
		builder.WriteRune(ch)
		l.advance()
	}
}

func (l *Lexer) skipWhitespace() {
	l.readWhile(unicode.IsSpace)
}

func (l *Lexer) readNumber() (string, error) {
	var builder strings.Builder

	dotReached := false

	for {
		ch := l.peek()
		if ch == 0 {
			break
		}

		if ch == '.' {
			if dotReached {
				return builder.String(), errors.New("double dot")
			}
			dotReached = true
		} else if !unicode.IsDigit(ch) {
			break
		}

		builder.WriteRune(ch)
		l.advance()
	}

	return builder.String(), nil
}

func (l *Lexer) readString() (string, error) {
	var builder strings.Builder

	builder.WriteRune('"')

	escaping := false

	l.advance()

	for {
		ch := l.peek()
		if ch == 0 {
			return builder.String(), errors.New("unterminated string")
		}

		l.advance()

		builder.WriteRune(ch)

		if escaping {
			escaping = false
		} else if ch == '\\' {
			escaping = true
		} else if ch == '"' {
			return builder.String(), nil
		}
	}
}

func (l *Lexer) readBitOrLogicalOperator(
	ch rune,
	bitSimpleType TokenType,
	bitAssignType TokenType,
	logicalSimpleType TokenType,
	logicalAssignType TokenType,
) (string, TokenType) {
	next := l.peek()

	if next != ch {
		return l.readOperator(ch, bitSimpleType, bitAssignType)
	}

	l.advance()

	value, tokenType := l.readOperator(next, logicalSimpleType, logicalAssignType)
	return string(ch) + value, tokenType
}

func (l *Lexer) readOperator(ch rune, simpleType TokenType, assignType TokenType) (string, TokenType) {
	next := l.peek()

	if next == '=' {
		l.advance()
		return string(ch) + string(next), assignType
	}

	return string(ch), simpleType
}

func (l *Lexer) nextToken() (Token, error) {
	l.skipWhitespace()
	ch := l.peek()

	token := Token{Column: l.column, Line: l.line, Pos: l.pos}

	if ch == 0 {
		token.Type = TokenEOF
		return token, nil
	}

	// number literals

	if unicode.IsDigit(ch) {
		var err error
		token.Value, err = l.readNumber()

		if strings.Contains(token.Value, ".") {
			token.Type = TokenFloatLiteral
		} else {
			token.Type = TokenIntLiteral
		}

		return token, err
	}

	// string literals

	if ch == '"' {
		token.Type = TokenStringLiteral

		var err error
		token.Value, err = l.readString()
		return token, err
	}

	// keywords & names

	if unicode.IsLetter(ch) || ch == '_' {
		token.Value = l.readWhile(func(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' })

		switch token.Value {
		case "function":
			token.Type = TokenFunction
		case "fragment":
			token.Type = TokenFragment
		case "vertex":
			token.Type = TokenVertex

		case "export":
			token.Type = TokenExport
		case "import":
			token.Type = TokenImport
		case "from":
			token.Type = TokenFrom

		case "let":
			token.Type = TokenLet
		case "return":
			token.Type = TokenReturn
		case "if":
			token.Type = TokenIf
		case "else":
			token.Type = TokenElse
		case "for":
			token.Type = TokenFor
		case "while":
			token.Type = TokenWhile
		case "do":
			token.Type = TokenDo
		case "uniform":
			token.Type = TokenUniform

		case "float":
			token.Type = TokenFloat
		case "int":
			token.Type = TokenInt
		case "bool":
			token.Type = TokenBoolean
		case "vec2":
			token.Type = TokenVec2
		case "vec3":
			token.Type = TokenVec3
		case "vec4":
			token.Type = TokenVec4
		case "ivec2":
			token.Type = TokenIntVec2
		case "ivec3":
			token.Type = TokenIntVec3
		case "ivec4":
			token.Type = TokenIntVec4
		case "bvec2":
			token.Type = TokenBooleanVec2
		case "bvec3":
			token.Type = TokenBooleanVec3
		case "bvec4":
			token.Type = TokenBooleanVec4

		case "true", "false":
			token.Type = TokenBooleanLiteral

		default:
			token.Type = TokenIdentifier
		}

		return token, nil
	}

	// single-character tokens, operators, comments

	l.advance()

	token.Value = string(ch)

	switch ch {
	case ';':
		token.Type = TokenSemicolon
	case ':':
		token.Type = TokenColon
	case ',':
		token.Type = TokenComma
	case '.':
		token.Type = TokenDot
	case '@':
		token.Type = TokenAt

	case '(':
		token.Type = TokenOpenParen
	case ')':
		token.Type = TokenCloseParen
	case '{':
		token.Type = TokenOpenBrace
	case '}':
		token.Type = TokenCloseBrace
	case '[':
		token.Type = TokenOpenBracket
	case ']':
		token.Type = TokenCloseBracket

	case '+':
		token.Value, token.Type = l.readOperator(ch, TokenPlus, TokenPlusAssign)
	case '-':
		token.Value, token.Type = l.readOperator(ch, TokenMinus, TokenMinusAssign)
	case '/':
		if l.peek() == '/' {
			token.Value, token.Type = "/"+l.readWhile(func(r rune) bool { return r != '\n' }), TokenComment
			l.advance()
			if !l.ParseComments {
				return l.nextToken()
			}
			break
		}
		token.Value, token.Type = l.readOperator(ch, TokenDivide, TokenDivideAssign)
	case '*':
		token.Value, token.Type = l.readOperator(ch, TokenMultiply, TokenMultiplyAssign)

	case '=':
		if l.peek() == '=' {
			token.Value, token.Type = "==", TokenEqual
			l.advance()
		} else {
			token.Type = TokenAssign
		}

	case '!':
		token.Value, token.Type = l.readOperator(ch, TokenNot, TokenNotEqual)
	case '^':
		token.Value, token.Type = l.readOperator(ch, TokenXor, TokenXorAssign)
	case '&':
		token.Value, token.Type = l.readBitOrLogicalOperator(
			ch,
			TokenBitAnd,
			TokenBitAndAssign,
			TokenLogicalAnd,
			TokenLogicalAndAssign,
		)
	case '|':
		token.Value, token.Type = l.readBitOrLogicalOperator(
			ch,
			TokenBitOr,
			TokenBitOrAssign,
			TokenLogicalOr,
			TokenLogicalOrAssign,
		)

	default:
		return token, fmt.Errorf("unknown token starting with %c", ch)
	}

	return token, nil
}

func (l *Lexer) Tokenize(input []rune) ([]Token, []*LexerError) {
	l.input = input

	l.line = 1
	l.column = 1
	l.pos = 0

	var errors []*LexerError
	result := make([]Token, 0)

	for {
		token, err := l.nextToken()
		if err != nil {
			lexerError := &LexerError{Token: &token, Err: err}

			if errors == nil {
				errors = []*LexerError{lexerError}
			} else {
				errors = append(errors, lexerError)
			}

			if !l.ParseAllErrors {
				return result, errors
			}
		}
		result = append(result, token)
		if token.Type == TokenEOF {
			return result, errors
		}
	}
}
