package driftc

type TokenType string

const (
	// general, done

	TokenEOF       TokenType = "EOF"
	TokenName      TokenType = "Name"
	TokenSemicolon TokenType = ";"
	TokenColon     TokenType = ":"
	TokenComma     TokenType = ","
	TokenDot       TokenType = "."
	TokenLet       TokenType = "let"
	TokenReturn    TokenType = "return"

	// functions, done

	TokenFunction TokenType = "function"
	TokenFragment TokenType = "fragment"
	TokenVertex   TokenType = "vertex"

	// imports, done

	TokenExport TokenType = "export"
	TokenImport TokenType = "import"
	TokenFrom   TokenType = "from"

	// types, done

	TokenFloat       TokenType = "float"
	TokenInt         TokenType = "int"
	TokenBoolean     TokenType = "bool"
	TokenVec2        TokenType = "vec2"
	TokenVec3        TokenType = "vec3"
	TokenVec4        TokenType = "vec4"
	TokenIntVec2     TokenType = "ivec2"
	TokenIntVec3     TokenType = "ivec3"
	TokenIntVec4     TokenType = "ivec4"
	TokenBooleanVec2 TokenType = "bvec2"
	TokenBooleanVec3 TokenType = "bvec3"
	TokenBooleanVec4 TokenType = "bvec4"

	// literals, done

	TokenStringLiteral  TokenType = "string literal"
	TokenFloatLiteral   TokenType = "float literal"
	TokenIntLiteral     TokenType = "int literal"
	TokenBooleanLiteral TokenType = "bool literal"

	// brackets, done

	TokenOpenParen    TokenType = "("
	TokenCloseParen   TokenType = ")"
	TokenOpenBrace    TokenType = "{"
	TokenCloseBrace   TokenType = "}"
	TokenOpenBracket  TokenType = "["
	TokenCloseBracket TokenType = "]"

	// general operators, done

	TokenPlus     TokenType = "+"
	TokenMinus    TokenType = "-"
	TokenDivide   TokenType = "/"
	TokenMultiply TokenType = "*"

	// logical & bit operators, done

	TokenEqual      TokenType = "=="
	TokenNot        TokenType = "!"
	TokenNotEqual   TokenType = "!="
	TokenXor        TokenType = "^"
	TokenBitAnd     TokenType = "&"
	TokenLogicalAnd TokenType = "&&"
	TokenBitOr      TokenType = "|"
	TokenLogicalOr  TokenType = "||"

	// assign operators, done

	TokenAssign           TokenType = "="
	TokenPlusAssign       TokenType = "+="
	TokenMinusAssign      TokenType = "-="
	TokenDivideAssign     TokenType = "/="
	TokenMultiplyAssign   TokenType = "*="
	TokenXorAssign        TokenType = "^="
	TokenBitOrAssign      TokenType = "|="
	TokenLogicalOrAssign  TokenType = "||="
	TokenBitAndAssign     TokenType = "&="
	TokenLogicalAndAssign TokenType = "&&="
)

type Token struct {
	Type  TokenType
	Value string

	Line   int
	Column int
	Pos    int
}
