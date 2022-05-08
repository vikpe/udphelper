package udphelper_test

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/udphelper"
)

func TestUdpServer_Listen(t *testing.T) {
	udpServer := udphelper.New(":8000")

	go func() {
		udpServer.Listen()
	}()
	time.Sleep(10 * time.Millisecond)

	response := sendUdpPacket(":8000", []byte("ping"))
	assert.Equal(t, []byte("ping"), udpServer.Requests[0])
	assert.Equal(t, []byte{}, response)
}

func TestUdpServer_Echo(t *testing.T) {
	udpServer := udphelper.New(":8001")

	go func() {
		udpServer.Echo()
	}()
	time.Sleep(10 * time.Millisecond)

	response := sendUdpPacket(":8001", []byte("ping"))
	assert.Equal(t, []byte("ping"), udpServer.Requests[0])
	assert.Equal(t, []byte("ok:ping"), response)
}

func TestUdpServer_Respond(t *testing.T) {
	t.Run("Single response", func(t *testing.T) {
		udpServer := udphelper.New(":8002")

		go func() {
			udpServer.Respond([]byte("pong"))
		}()
		time.Sleep(30 * time.Millisecond)

		response := sendUdpPacket(":8002", []byte("ping"))
		assert.Equal(t, []byte("ping"), udpServer.Requests[0])
		assert.Equal(t, []byte("pong"), response)
	})

	t.Run("Multiple responses", func(t *testing.T) {
		udpServer := udphelper.New(":8003")

		go func() {
			udpServer.Respond([]byte("pong"), []byte("beta"))
		}()
		time.Sleep(30 * time.Millisecond)

		response1 := sendUdpPacket(":8003", []byte("ping"))
		response2 := sendUdpPacket(":8003", []byte("alpha"))

		assert.Equal(t, []byte("ping"), udpServer.Requests[0])
		assert.Equal(t, []byte("pong"), response1)

		assert.Equal(t, []byte("alpha"), udpServer.Requests[1])
		assert.Equal(t, []byte("beta"), response2)
	})
}

func sendUdpPacket(addr string, packet []byte) []byte {
	conn, _ := net.Dial("udp4", addr)
	defer conn.Close()
	timeoutInMs := 30

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
