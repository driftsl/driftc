package driftc

type RootNode struct {
	Imports []ImportNode
}

type ImportNode struct {
	To   any   // Must be identifier token or DeconstructionNode
	From Token // Must be string literal token
}

type DeconstructionNode struct {
	NameTokens []Token // Must be identifier tokens
}
