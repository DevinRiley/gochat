package main

import (
  "fmt"
  "net"
  "bufio"
  "strings"
)


func getUsername(conn net.Conn) string {
  conn.Write([]byte("Enter your user name: "))
  username, _ := bufio.NewReader(conn).ReadString('\n')

  return strings.Trim(username, "\n")
}

func handleConnection(conn net.Conn, broadcast func(string, string)) {
  username := getUsername(conn)

  for {
      message, _ := bufio.NewReader(conn).ReadString('\n')
      fmt.Println("In handleConnection")
      broadcast(username, message)
  }
}

func main() {
  ln, _ := net.Listen("tcp", ":8001")
  connpool := make([]net.Conn, 1, 100)

  broadcast := func(name string, s string) {
    for _, conn := range connpool {
      if conn != nil {
        conn.Write([]byte(name + ": " + s + "\n"))
      }
    }
  }

  for {
    conn, _ := ln.Accept()
    connpool = append(connpool, conn)
    go handleConnection(conn, broadcast)
  }
}
