package server

import (
	"bufio"
	"crypto/tls"
	"io"
	"net"
)

func (s *Server) httpsHandler(conn net.Conn) {
	s.httpHandler(tls.Server(conn, s.TLSConfig))
}

func (s *Server) httpHandler(conn net.Conn) {
	s.socksHandler(newHTTPStripper(s, conn))
}

type httpStripper struct {
	net.Conn
	server     *Server
	body       *io.ReadCloser
	sentHeader bool
	ioBuf      *bufio.Reader
}

func newHTTPStripper(server *Server, conn net.Conn) *httpStripper {
	return &httpStripper{
		Conn:   conn,
		server: server,
		ioBuf:  bufio.NewReader(conn),
	}
}
