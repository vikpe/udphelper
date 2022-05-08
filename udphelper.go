package udphelper

import (
	"net"
)

type UdpServer struct {
	Address   string
	Requests  [][]byte
	conn      net.PacketConn
	echo      bool
	responses [][]byte
}

func New(address string) *UdpServer {
	server := UdpServer{
		Address:   address,
		Requests:  make([][]byte, 0),
		echo:      false,
		responses: make([][]byte, 0),
	}
	return &server
}

func (s *UdpServer) Listen() {
	s.conn, _ = net.ListenPacket("udp", s.Address)
	defer s.conn.Close()

	responseCount := len(s.responses)

	for {
		responseIndex := len(s.Requests)
		buffer := make([]byte, 1024)
		packetLength, dst, _ := s.conn.ReadFrom(buffer)
		packet := buffer[:packetLength]

		if s.echo || responseCount > 0 {
			var reply []byte = nil

			if s.echo {
				reply = append([]byte("ok:"), packet...)
			} else {
				reply = s.responses[responseIndex]
			}

			s.conn.WriteTo(reply, dst)
		}

		s.Requests = append(s.Requests, packet)
		isDone := responseCount > 1 && len(s.Requests) == responseCount

		if isDone {
			break
		}
	}
}

func (s *UdpServer) Echo() {
	s.echo = true
	s.Listen()
}

func (s *UdpServer) Respond(responses ...[]byte) {
	s.responses = append(s.responses, responses...)
	s.Listen()
}
