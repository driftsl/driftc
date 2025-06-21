package driftc

type TokenType int

const (
	// general, done

	TokenEOF TokenType = iota
	TokenName
	TokenSemicolon // ;
	TokenColon     // :
	TokenComma     // ,
	TokenDot       // .
	TokenLet       // let
	TokenReturn    // return

	// functions, done

	TokenFunction // function
	TokenFragment // fragment
	TokenVertex   // vertex

	// imports, done

	TokenExport // export
	TokenImport // import
	TokenFrom   // from

	// types, done

	TokenFloat       // float
	TokenInt         // int
	TokenBoolean     // bool
	TokenVec2        // vec2
	TokenVec3        // vec3
	TokenVec4        // vec4
	TokenIntVec2     // ivec2
	TokenIntVec3     // ivec3
	TokenIntVec4     // ivec4
	TokenBooleanVec2 // bvec2
	TokenBooleanVec3 // bvec3
	TokenBooleanVec4 // bvec4

	// literals, done

	TokenStringLiteral
	TokenFloatLiteral
	TokenIntLiteral
	TokenBooleanLiteral // true/false

	// brackets, done

	TokenOpenParen    // (
	TokenCloseParen   // )
	TokenOpenBrace    // {
	TokenCloseBrace   // }
	TokenOpenBracket  // [
	TokenCloseBracket // ]

	// general operators, done

	TokenPlus     // +
	TokenMinus    // -
	TokenDivide   // /
	TokenMultiply // *

	// logical & bit operators

	TokenEqual      // ==
	TokenNot        // !
	TokenNotEqual   // !=
	TokenXor        // ^
	TokenBitAnd     // &
	TokenLogicalAnd // &&
	TokenBitOr      // |
	TokenLogicalOr  // ||

	// assign operators (done)

	TokenAssign         // =
	TokenPlusAssign     // +=
	TokenMinusAssign    // -=
	TokenDivideAssign   // /=
	TokenMultiplyAssign // *=
	TokenBitOrAssign    // |=

	// not done

	TokenLogicalOrAssign  // ||=
	TokenBitAndAssign     // &=
	TokenLogicalAndAssign // &&=
	TokenXorAssign        // ^=
)

type Token struct {
	Type  TokenType
	Value string

	Line   int
	Column int
	Pos    int
}
