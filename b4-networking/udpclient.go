package main

import (
    "fmt"
    "net"
    "bufio"
)

func main() {
    // connect to UDP server
    p :=  make([]byte, 2048)
    conn, err := net.Dial("udp", "127.0.0.1:1234")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }

    // send message
    fmt.Fprintf(conn, "abc.test")
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }

    // close UDP connection
    conn.Close()
}
