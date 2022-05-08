package udph

import "net"

type UdpServer struct {
	Address        string
	conn           net.PacketConn
	responsePrefix string
}

func New(address string) *UdpServer {
	server := UdpServer{Address: address}
	return &server
}

func (s UdpServer) Listen() *UdpServer {
	s.conn, _ = net.ListenPacket("udp", s.Address)
	defer s.conn.Close()
	return &s
}

func (s UdpServer) Echo(prefix string) *UdpServer {
	s.responsePrefix = prefix
	return &s
}

/*
func udpListenAndEcho(addr string) {
	conn, _ := net.ListenPacket("udp", addr)
	buffer := make([]byte, 1024)
	messageLength, dst, _ := conn.ReadFrom(buffer)
	message := buffer[:messageLength]
	response := "OK: " + string(message)
	conn.WriteTo([]byte(response), dst)
}
*/
