package main

import (
  "fmt"
  "log"
  "net"
  "bufio"
)

func main() {
  count := 1000
  c := make(chan string)

  for i := 0; i < count; i++ {
    go func(c chan string) {
      conn, err := net.Dial("tcp", "localhost:2000")
      if err != nil {
        log.Fatal(err)
      }

      fmt.Fprintf(conn, "Hello!\n")
      response, err := bufio.NewReader(conn).ReadString('\n')
      conn.Close()
      c <- response
    }(c)
  }

  for i := 0; i < count; i++ {
    log.Printf("Got back: %s\n", <-c)
  }
}
