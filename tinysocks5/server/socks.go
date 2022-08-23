package server

import "net"

func (s *Server) socksHandler(cn net.Conn) {
	defer cn.Close()

}
