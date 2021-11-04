package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type StatsConn struct {
	net.Conn

	bytesRead uint64
}

func NewStatsConn(conn net.Conn) *StatsConn {
	return &StatsConn{
		Conn:      conn,
		bytesRead: 0,
	}
}

func (sc *StatsConn) BytesRead() uint64 {
	return sc.bytesRead
}

func (sc *StatsConn) Read(p []byte) (int, error) {
	n, err := sc.Conn.Read(p)
	sc.bytesRead += uint64(n)
	return n, err
}

func printAll(conn net.Conn) {
	numBytes, err := fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytes written %v\n", numBytes)
	response, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response = %v\n", string(response[:50]))
}

func regular() {
	conn, err := net.Dial("tcp", "golang.org:http")
	if err != nil {
		log.Fatal(err)
	}
	printAll(conn)
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func withStats() {
	conn, err := net.Dial("tcp", "golang.org:http")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()
	statsConn := NewStatsConn(conn)
	printAll(statsConn)
	fmt.Printf("Bytes read was %v\n", statsConn.BytesRead())
}

func main() {
	regular()
	withStats()
}
