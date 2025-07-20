package driftls

import (
	"io"
	"strconv"
	"strings"
)

func (s *Server) nextRequest() ([]byte, error) {
	var contentLength int

	for {
		line, err := s.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = trimLineBreak(line)

		if line == "" {
			break
		}

		headerParts := strings.SplitN(line, ": ", 2)

		switch strings.ToLower(headerParts[0]) {
		case "content-length":
			contentLength, err = strconv.Atoi(headerParts[1])
			if err != nil {
				return nil, err
			}
		}
	}

	buf := make([]byte, contentLength)
	_, err := io.ReadFull(s.reader, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func trimLineBreak(line string) string {
	length := len(line)
	if length >= 2 && line[length-2] == '\r' {
		return line[:length-2]
	} else {
		return line[:length-1]
	}
}
