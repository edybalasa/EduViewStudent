package networking

import (
	"fmt"
	"net"
	"os"
	"time"
)

type SignalSender struct {
	hostname string
}

func (sg SignalSender) PrepareSendService() {
	broadcastAddr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:25643")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp4", nil, broadcastAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	sg.hostname, err = os.Hostname()
	if err != nil {
		fmt.Println("Error retreiving hostname: ", err)
	}

	message := []byte(sg.hostname)

	sg.runPairingService(conn, message)
}

func (sg SignalSender) PrepareReceivingService() net.Conn {
	addr, err := net.ResolveUDPAddr("udp4", ":25643")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Listening for broadcast messages on port 25643...")

	return conn
}

func (sg SignalSender) runPairingService(sendConn net.Conn, message []byte) {
	for true {
		_, err := sendConn.Write(message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			os.Exit(1)
		}

		fmt.Println("Broadcast byte sent!")
		time.Sleep(5000 * time.Millisecond)
	}
}

func (sg SignalSender) SendSignals() {
	sg.PrepareSendService()
}
