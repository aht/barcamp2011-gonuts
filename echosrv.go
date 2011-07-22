package main

import ("net"; "io"; "os"; "log")

func die(err os.Error) {
    if err != nil { log.Fatal(err) }
}

func main() {
    listener, err := net.Listen("tcp", "127.0.0.1:3640"); die(err)
    for {
        conn, err := listener.Accept(); die(err)
        go func() {
            defer conn.Close()
            n, err := io.Copy(conn, conn); die(err)
            log.Printf("echoed %d byte from %s\n", n, conn.RemoteAddr())
        }()
    }
}
