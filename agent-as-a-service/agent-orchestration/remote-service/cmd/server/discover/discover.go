package discover

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
)

const (
	discoveryPort = 9999
	bufferSize    = 1024
)

type DeviceInfo struct {
	Name       string `json:"name"`
	IP         string `json:"ip"`
	OS         string `json:"os"`
	Hostname   string `json:"hostname"`
	MACAddress string `json:"mac_address"`
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("Error getting local IP: %v", err)
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getMACAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Error getting MAC address: %v", err)
		return ""
	}

	for _, i := range interfaces {
		if i.Flags&net.FlagUp != 0 && !bytes.Equal(i.HardwareAddr, nil) {
			return i.HardwareAddr.String()
		}
	}
	return ""
}

func DiscoverStart() {
	// Get device information
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error getting hostname: %v", err)
		hostname = "unknown"
	}

	deviceInfo := DeviceInfo{
		Name:       "Neuron Device",
		IP:         getLocalIP(),
		OS:         runtime.GOOS,
		Hostname:   hostname,
		MACAddress: getMACAddress(),
	}

	// Create UDP address
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", discoveryPort))
	if err != nil {
		log.Fatalf("Error resolving UDP address: %v", err)
	}

	// Create UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Error creating UDP connection: %v", err)
	}
	defer conn.Close()

	log.Printf("Device discovery server started on port %d", discoveryPort)
	log.Printf("Device info: %+v", deviceInfo)

	buffer := make([]byte, bufferSize)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading UDP: %v", err)
			continue
		}

		message := string(buffer[:n])
		log.Printf("Received message from %s: %s", remoteAddr, message)

		if message == "AGENT_DISCOVER" {
			// Marshal device info to JSON
			response, err := json.Marshal(deviceInfo)
			if err != nil {
				log.Printf("Error marshaling response: %v", err)
				continue
			}

			// Send response
			_, err = conn.WriteToUDP(response, remoteAddr)
			if err != nil {
				log.Printf("Error sending response: %v", err)
				continue
			}

			log.Printf("Sent device info to %s", remoteAddr)
		}
	}
}
