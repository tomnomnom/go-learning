package main

import (
  "log"
  "path/filepath"
  "time"
)

func main() {

  for {
    files, err := filepath.Glob("*.log")
    if err != nil {
      log.Fatal(err)
    }
    if len(files) == 0 {
      time.Sleep(1 * time.Second)
      continue
    }
    log.Printf("%v\n", files)
    time.Sleep(1 * time.Second)
  }
}
