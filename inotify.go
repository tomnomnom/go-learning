package main

import (
  "syscall"
  "fmt"
  "log"
)

func main() {
  fd, err := syscall.InotifyInit()
  if err != nil {
    log.Fatal(err)
  }
  defer syscall.Close(fd)

  wd, err := syscall.InotifyAddWatch(fd, "test1.log", syscall.IN_ALL_EVENTS)
  if err != nil {
    log.Fatal(err)
  }
  defer syscall.InotifyRmWatch(fd, uint32(wd))

  event := make([]byte, syscall.SizeofInotifyEvent)
  for {
    _, err := syscall.Read(fd, event)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%#v\n", event)
    //fmt.Printf("%#V\n", syscall.IN_ALL_EVENTS)
  }

}
