package driftls

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	reader *bufio.Reader
	writer *bufio.Writer

	alive bool

	documents DocumentsVault
}

type TextDocumentIdentifier struct {
	Uri string `json:"uri"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

func NewServer(r *bufio.Reader, w *bufio.Writer) *Server {
	return &Server{
		reader: r,
		writer: w,

		documents: DocumentsVault{
			Documents: make(map[string]string),
		},
	}
}

func (s *Server) Serve() error {
	s.alive = true

	for s.alive {
		body, err := s.nextRequest()
		if err != nil {
			return err
		}

		var data JsonRpcRequest[json.RawMessage]
		if err := json.Unmarshal(body, &data); err != nil {
			return err
		}

		fmt.Fprintln(os.Stderr, data.Method, data.Id)

		switch data.Method {
		case "initialize":
			s.initialize(data.Id)

		case "textDocument/didOpen":
			if err := s.documents.Open(data.Params); err != nil {
				return err
			}
		case "textDocument/didChange":
			if err := s.documents.Change(data.Params); err != nil {
				return err
			}
		case "textDocument/didClose":
			if err := s.documents.Close(data.Params); err != nil {
				return err
			}
		case "textDocument/semanticTokens/full":
			if err := s.sendTokens(data.Id, data.Params); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Server) sendServerResponse(id any, r any) error {
	return s.sendJsonRpcResponse(id, r, nil)
}

func (s *Server) sendServerError(id any, code int, message string) error {
	return s.sendJsonRpcResponse(id, nil, &JsonRpcError{Code: code, Message: message})
}

func (s *Server) sendNotification(method string, params any) error {
	return s.sendJson(JsonRpcRequest[any]{
		JsonRpc: "2.0",
		Method:  method,
		Params:  params,
	})
}

func (s *Server) sendJsonRpcResponse(id any, result any, jsonRpcError *JsonRpcError) error {
	return s.sendJson(JsonRpcResponse{
		JsonRpc: "2.0",
		Id:      id,
		Result:  result,
		Error:   jsonRpcError,
	})
}

func (s *Server) sendJson(object any) error {
	data, err := json.Marshal(object)
	if err != nil {
		return err
	}

	return s.send(data)
}

func (s *Server) send(data []byte) error {
	if _, err := s.writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n", len(data))); err != nil {
		return err
	}

	if _, err := s.writer.Write(data); err != nil {
		return err
	}

	return s.writer.Flush()
}
