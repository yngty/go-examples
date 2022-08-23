package socks

import (
	"errors"
	"io"
	"net"
	"strconv"
)

const (
	Version     = 5
	UserPassVer = 1
)

const (
	MethodNoAuth uint8 = iota
	MethodGSSAPI
	MethodUserPass
	MethodNoAcceptable uint8 = 0xFF
)

const (
	CmdConnect uint8 = iota + 1
	CmdBind
	CmdUDP
	CmdUDPOverTCP
)

const (
	AddrIPv4   uint8 = 1
	AddrDomain       = 3
	AddrIPv6         = 4
)

const (
	Succeeded uint8 = iota
	Failure
	Allowed
	NetUnreachable
	HostUnreachable
	ConnRefused
	TTLExpired
	CmdUnsupported
	AddrUnsupported
)

var (
	ErrBadVersion  = errors.New("bad version")
	ErrBadFormat   = errors.New("bad format")
	ErrBadAddrType = errors.New("bad address type")
	ErrShortBuffer = errors.New("short buffer")
	ErrBadMethod   = errors.New("bad method")
	ErrAuthFailure = errors.New("auth failure")
)

func WriteMethod(method uint8, w io.Writer) error {
	_, err := w.Write([]byte{Version, method})
	return err
}

func WriteMethods(methods []uint8, w io.Writer) error {
	b := make([]byte, 2+len(methods))
	b[0] = Version
	b[1] = uint8(len(methods))
	copy(b[2:], methods)
	_, err := w.Write(b)
	return err
}

type Addr struct {
	Type uint8
	Host string
	Port uint16
}

func NewAddr(sa string) (addr *Addr, error) {
	host, port, err := net.SplitHostPort(sa)
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	addr = NewAddrFromPair(host, port)
	return
}

func NewAddrFromPair(host string, port int) (addr *Addr) {
	addr = &Addr{
		Type: AddrDomain,
		Host: host,
		Port: uint16(port),
	}

	if ip := net.ParseIP(host); ip != nil {
		if ip.To4() != nil {
			addr.Type = AddrIPv4
		} else {
			addr.Type = AddrIPv6
		}
	}

	return
}

func NewFromAddr(ln, conn net.Addr) (addr *Addr, err error)  {
	_, sport, err := net.SplitHostPort(ln.String())
	if err != nil {
		return nil, err
	}

	host, _, err := net.SplitHostPort(con.String())
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(sport)
	if err != nil {
		return nil, err
	}

	addr = NewAddrFromPair(host, port)
	return
}


type Request struct {
	Cmd  uint8
	Addr *Addr
}

func newRequest(cmd uint8, addr *Addr) *Request {
	return &Request{
		Cmd:  cmd,
		Addr: addr,
	}

}
