package main

import (
  "os"
  "bufio"
  "fmt"
  "time"
)

func followFile(filename string, lines chan string) {
  fd, err := os.Open(filename)
  if err != nil {
    return
  }
  defer fd.Close()

  // Seek to the end
  _, err = fd.Seek(0, os.SEEK_END)
  if err != nil {
    return
  }

  r := bufio.NewReader(fd)

  for {
    line, err := r.ReadString('\n')
    if err != nil {
      time.Sleep(100 * time.Millisecond)
      continue
    }

    lines <- line
  }
}

func main() {
  lines := make(chan string)
  go followFile("test1.log", lines)
  go followFile("test2.log", lines)

  for {
    fmt.Printf("%s", <-lines)
  }
}
