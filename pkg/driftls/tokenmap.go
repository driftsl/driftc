package driftls

import (
	"fmt"

	"github.com/driftsl/driftc/pkg/driftc"
)

var tokensArray = [...]string{
	"keyword",
	"type",
	"variable",
	"number",
	"string",
	"comment",
	"operator",
}

func mapTokenType(t driftc.TokenType) int {
	switch t {
	// skip
	case driftc.TokenEOF,
		driftc.TokenColon,
		driftc.TokenSemicolon,
		driftc.TokenDot,
		driftc.TokenComma,
		driftc.TokenOpenBrace,
		driftc.TokenOpenBracket,
		driftc.TokenOpenParen,
		driftc.TokenCloseBrace,
		driftc.TokenCloseBracket,
		driftc.TokenCloseParen:
		return -1

	// keywords
	case driftc.TokenLet,
		driftc.TokenFunction,
		driftc.TokenReturn,
		driftc.TokenImport,
		driftc.TokenExport,
		driftc.TokenFrom,
		driftc.TokenVertex,
		driftc.TokenFragment:
		return 0

	// types
	case driftc.TokenBoolean,
		driftc.TokenFloat,
		driftc.TokenInt,
		driftc.TokenVec2,
		driftc.TokenVec3,
		driftc.TokenVec4,
		driftc.TokenIntVec2,
		driftc.TokenIntVec3,
		driftc.TokenIntVec4,
		driftc.TokenBooleanVec2,
		driftc.TokenBooleanVec3,
		driftc.TokenBooleanVec4:
		return 1

	// names
	case driftc.TokenName:
		return 2

	// literals
	case driftc.TokenFloatLiteral,
		driftc.TokenIntLiteral,
		driftc.TokenBooleanLiteral:
		return 3
	case driftc.TokenStringLiteral:
		return 4

	// comments
	case driftc.TokenComment:
		return 5

	// operators
	case driftc.TokenPlus,
		driftc.TokenMinus,
		driftc.TokenDivide,
		driftc.TokenMultiply,
		driftc.TokenEqual,
		driftc.TokenNot,
		driftc.TokenNotEqual,
		driftc.TokenXor,
		driftc.TokenBitAnd,
		driftc.TokenLogicalAnd,
		driftc.TokenBitOr,
		driftc.TokenLogicalOr,
		driftc.TokenAssign,
		driftc.TokenPlusAssign,
		driftc.TokenMinusAssign,
		driftc.TokenDivideAssign,
		driftc.TokenMultiplyAssign,
		driftc.TokenXorAssign,
		driftc.TokenBitOrAssign,
		driftc.TokenLogicalOrAssign,
		driftc.TokenBitAndAssign,
		driftc.TokenLogicalAndAssign:
		return 6
	}
	panic(fmt.Sprintf("unexpected driftc.TokenType: %#v", t))
}
