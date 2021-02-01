package main
import (
  "flag"
  "fmt"
  "net"
  "os"
)
var host = flag.String("host", "", "host")
var port = flag.String("port", "5060", "port")
func main() {
  flag.Parse()
  addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
  if err != nil {
    fmt.Println("Can't resolve address: ", err)
    os.Exit(1)
  }
  conn, err := net.ListenUDP("udp", addr)
  if err != nil {
    fmt.Println("Error listening:", err)
    os.Exit(1)
  }
  defer conn.Close()
  for {
    handleClient(conn)
  }
}
func handleClient(conn *net.UDPConn) {
  fmt.Println("------------------------------------")
  data := make([]byte, 1024)
  n, remoteAddr, err := conn.ReadFromUDP(data)
  if err != nil {
    fmt.Println("failed to read UDP msg because of ", err.Error())
    return
  }
  fmt.Println("client addr: " + remoteAddr.String())
  fmt.Println(n, remoteAddr)
  conn.WriteToUDP(data[:n], remoteAddr)
}
