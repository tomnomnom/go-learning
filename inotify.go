package main

import (
  "fmt"
  "log"
  "syscall"
  "unsafe"
)

func main() {
  fd, err := syscall.InotifyInit()
  if err != nil {
    log.Fatal(err)
  }
  defer syscall.Close(fd)

  wd, err := syscall.InotifyAddWatch(fd, "test1.log", syscall.IN_ALL_EVENTS)
  _, err = syscall.InotifyAddWatch(fd, "../test2.log", syscall.IN_ALL_EVENTS)
  //_, err = syscall.InotifyAddWatch(fd, ".", syscall.IN_ALL_EVENTS)
  if err != nil {
    log.Fatal(err)
  }
  defer syscall.InotifyRmWatch(fd, uint32(wd))

  fmt.Printf("WD is %d\n", wd)

  for {
    // Room for at least 128 events
    buffer := make([]byte, syscall.SizeofInotifyEvent*128)
    bytesRead, err := syscall.Read(fd, buffer)
    if err != nil {
      log.Fatal(err)
    }

    if bytesRead < syscall.SizeofInotifyEvent {
      // No point trying if we don't have at least one event
      continue
    }

    fmt.Printf("Size of InotifyEvent is %s\n", syscall.SizeofInotifyEvent)
    fmt.Printf("Bytes read: %d\n", bytesRead)

    offset := 0
    for offset < bytesRead-syscall.SizeofInotifyEvent {
      event := (*syscall.InotifyEvent)(unsafe.Pointer(&buffer[offset]))
      fmt.Printf("%+v\n", event)

      if (event.Mask & syscall.IN_ACCESS) > 0 {
        fmt.Printf("Saw IN_ACCESS for %+v\n", event)
      }

      // We need to account for the length of the name
      offset += syscall.SizeofInotifyEvent + int(event.Len)
    }
  }

}
