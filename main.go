package main

import (
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":10250")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		/*
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				io.WriteString(conn, scanner.Text())
			}
		*/
		go func() {
			io.Copy(conn, conn)
			conn.Close()
		}()
	}
}
