package main

import (
      "fmt"
      "os"
      "os/signal"
)

func main() {
      c := make(chan string)
      d := make(chan string)

      go func() {
          for {
              c <- "value from goroutine" // blocks until there is a receiver
          }
      }()
      go func() {
          for {
              d <- ("processed " + <-c) // blocks on receiving, then block on sending
          }
      }()

      for {
          select {
          case val := <-c:
              fmt.Println(val)
          case val := <-d:
              fmt.Println(val)
          case sig := <-signal.Incoming:
              fmt.Println("got signal:", sig)
              os.Exit(0)
          }
      }
}
