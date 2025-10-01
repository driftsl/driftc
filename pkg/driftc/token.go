package driftc

type TokenType string

const (
	// general

	TokenUnknown    TokenType = ""
	TokenEOF        TokenType = "EOF"
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenSemicolon  TokenType = ";"
	TokenColon      TokenType = ":"
	TokenComma      TokenType = ","
	TokenDot        TokenType = "."
	TokenComment    TokenType = "COMMENT"
	TokenAt         TokenType = "@"

	// functions

	TokenFunction TokenType = "function"
	TokenFragment TokenType = "fragment"
	TokenVertex   TokenType = "vertex"

	// imports

	TokenExport TokenType = "export"
	TokenImport TokenType = "import"
	TokenFrom   TokenType = "from"

	// other keywords

	TokenLet     TokenType = "let"
	TokenReturn  TokenType = "return"
	TokenIf      TokenType = "if"
	TokenElse    TokenType = "else"
	TokenFor     TokenType = "for"
	TokenWhile   TokenType = "while"
	TokenDo      TokenType = "do"
	TokenUniform TokenType = "uniform"

	// types

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

	// literals

	TokenStringLiteral  TokenType = "STRING LITERAL"
	TokenFloatLiteral   TokenType = "FLOAT LITERAL"
	TokenIntLiteral     TokenType = "INT LITERAL"
	TokenBooleanLiteral TokenType = "BOOLEAN LITERAL"

	// brackets

	TokenOpenParen    TokenType = "("
	TokenCloseParen   TokenType = ")"
	TokenOpenBrace    TokenType = "{"
	TokenCloseBrace   TokenType = "}"
	TokenOpenBracket  TokenType = "["
	TokenCloseBracket TokenType = "]"

	// general operators

	TokenPlus     TokenType = "+"
	TokenMinus    TokenType = "-"
	TokenDivide   TokenType = "/"
	TokenMultiply TokenType = "*"

	// logical & bit operators

	TokenEqual      TokenType = "=="
	TokenNot        TokenType = "!"
	TokenNotEqual   TokenType = "!="
	TokenXor        TokenType = "^"
	TokenBitAnd     TokenType = "&"
	TokenLogicalAnd TokenType = "&&"
	TokenBitOr      TokenType = "|"
	TokenLogicalOr  TokenType = "||"

	// assign operators

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
