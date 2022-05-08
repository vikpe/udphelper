package udphelper_test

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/udphelper"
)

func TestUdpServer_Listen(t *testing.T) {
	var udpServer udphelper.UdpServer

	go func() {
		udpServer = udphelper.New(":8000").Listen().Listen()
	}()
	time.Sleep(10 * time.Millisecond)

	response := sendUdpPacket(":8000", []byte("ping"))
	assert.Equal(t, []byte("ping"), udpServer.Requests[0])
	assert.Equal(t, []byte{}, response)
}

func TestUdpServer_Echo(t *testing.T) {
	var udpServer udphelper.UdpServer

	go func() {
		udpServer = udphelper.New(":8001").Echo()
	}()
	time.Sleep(10 * time.Millisecond)

	response := sendUdpPacket(":8001", []byte("ping"))
	assert.Equal(t, []byte("ping"), udpServer.Requests[0])
	assert.Equal(t, []byte("ok:ping"), response)
}

func TestUdpServer_Respond(t *testing.T) {
	var udpServer udphelper.UdpServer

	go func() {
		udpServer = udphelper.New(":8002").Respond([]byte("pong"))
	}()
	time.Sleep(10 * time.Millisecond)

	response := sendUdpPacket(":8002", []byte("ping"))
	assert.Equal(t, []byte("ping"), udpServer.Requests[0])
	assert.Equal(t, []byte("pong"), response)
}

func sendUdpPacket(addr string, packet []byte) []byte {
	conn, _ := net.Dial("udp4", addr)
	defer conn.Close()
	timeoutInMs := 50

	conn.SetDeadline(getDeadline(timeoutInMs))
	conn.Write(packet)

	responseBuffer := make([]byte, 8192)
	responseLength := 0
	conn.SetDeadline(getDeadline(timeoutInMs))
	responseLength, _ = conn.Read(responseBuffer)

	return responseBuffer[:responseLength]
}

func getDeadline(timeoutInMs int) time.Time {
	return time.Now().Add(time.Duration(timeoutInMs) * time.Millisecond)
}
