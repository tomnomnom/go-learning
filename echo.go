package main

import (
  "io"
  "log"
  "net"
)

func main() {
  l, err := net.Listen("tcp", ":2000")
  if err != nil {
    log.Fatal(err)
  }

  for {
    log.Println("Waiting for connection")
    conn, err := l.Accept()
    log.Printf("Accepted connection from %s\n", conn.RemoteAddr())
    if err != nil {
      log.Fatal(err)
    }

    go func(c net.Conn) {
      io.Copy(c, c)
      c.Close()
      log.Printf("Connection to %s closed\n", c.RemoteAddr())
    }(conn)
  }
}
