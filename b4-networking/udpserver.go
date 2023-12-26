package main

import (
    "fmt" 
    "net"
	"strings"
)


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, ip string) {
    _, err := conn.WriteToUDP([]byte(ip), addr)
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}

var dnsDomains = map[string]string{
	"abc.test": "192.168.10.1",
	"test.com": "192.168.10.2",
	"test.vn" : "192.168.10.3",
}


func main() {
	// setup UDP server
    p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 1234,
        IP: net.ParseIP("127.0.0.1"),
    }

    ser, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }

	// listening to client
    for {
        _, remoteaddr, err := ser.ReadFromUDP(p)

        fmt.Printf("Read a message from %v %s \n", remoteaddr, p)

        if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }

		// find ip from domain requested client
		domain := string(p)
		ip := ""
		for k, v := range dnsDomains {
			if strings.Contains(domain, k) {
				ip = v
				break
			}
		}

		// return ip
        go sendResponse(ser, remoteaddr, ip)
    }
}