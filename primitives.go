package main

func main() {
      c := make(chan string)
      d := make(chan string)

      go func() {
          for {
              c <- "msg" // sending blocks until there is a receiver
          }
      }()
      go func() {
          for {
              d <- ("relayed " + <-c) // process received value, then send it out
          }
      }()

      for {
          select {
          case val := <-c:
              println(val)
          case val := <-d:
              println(val)
          }
      }
}
