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

func (sg SignalSender) SendSignals() {
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

	for true {
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			os.Exit(1)
		}

		fmt.Println("Broadcast byte sent!")
		time.Sleep(5000 * time.Millisecond)
	}
}
