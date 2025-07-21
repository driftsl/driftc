package driftls

type JsonRpcRequest[T any] struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  T      `json:"params"`
	Id      any    `json:"id,omitempty"`
}

type JsonRpcResponse struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      any           `json:"id"`
	Result  any           `json:"result,omitempty"`
	Error   *JsonRpcError `json:"error,omitempty"`
}

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	JsonRpcParseErrorCode     = -32700 // Invalid JSON was received by the server.
	JsonRpcInvalidRequestCode = -32600 // The JSON sent is not a valid Request object.
	JsonRpcMethodNotFoundCode = -32601 // The method does not exist / is not available.
	JsonRpcInvalidParamsCode  = -32602 // Invalid method parameter(s).
	JsonRpcInternalErrorCode  = -32603 // Internal JSON-RPC error.
	// -32000 to -32099 implementation-defined server-errors
)
