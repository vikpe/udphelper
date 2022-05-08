package udphelper

import (
	"net"
)

type UdpServer struct {
	Address  string
	Requests [][]byte
	conn     net.PacketConn
	echo     bool
	response []byte
}

func New(address string) UdpServer {
	server := UdpServer{
		Address:  address,
		Requests: make([][]byte, 0),
		echo:     false,
	}
	return server
}

func (s UdpServer) Listen() UdpServer {
	if s.conn != nil {
		return s
	}

	s.conn, _ = net.ListenPacket("udp", s.Address)

	buffer := make([]byte, 1024)
	messageLength, dst, _ := s.conn.ReadFrom(buffer)
	message := buffer[:messageLength]
	s.Requests = append(s.Requests, message)

	if s.echo || s.response != nil {
		var reply []byte

		if s.echo {
			reply = append([]byte("ok:"), message...)

		} else {
			reply = s.response
		}

		s.conn.WriteTo(reply, dst)
	}

	return s
}

func (s UdpServer) Echo() UdpServer {
	s.echo = true
	return s.Listen()
}

func (s UdpServer) Respond(message []byte) UdpServer {
	s.response = message
	return s.Listen()
}
