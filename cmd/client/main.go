package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr := os.Getenv("TCP_ADDR")
	if addr == "" {
		addr = ":9000"
	}

	conn, err := net.Dial("tcp", "localhost"+addr)
	if err != nil {
		panic("failed to connect to the server: " + err.Error())
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("failed to close connection:", err)
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read from buffer:", err)
				return
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err := fmt.Fprintln(conn, scanner.Text())
		if err != nil {
			fmt.Println("failed to show message from application:", err)
		}
	}
}
