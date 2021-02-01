package main
import (
  "flag"
  "fmt"
  "net"
  "os"
  "time"
  "strconv"
)
var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "50601", "port")
var progName = flag.String("name", "c1", "name")
//go run timeclient.go -host time.nist.gov
func main() {
  flag.Parse()
  addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
  if err != nil {
    fmt.Println("Can't resolve address: ", err)
    os.Exit(1)
  }
  conn, err := net.DialUDP("udp", nil, addr)
  if err != nil {
    fmt.Println("Can't dial: ", err)
    os.Exit(1)
  }
  defer conn.Close()
  // Set initial deadline.
  // conn.SetReadDeadline(time.Now().Add(time.Minute))

  num := 0
  for {
    _, err = conn.Write([]byte(*progName + strconv.Itoa(num)))
    if err != nil {
      fmt.Println("failed:", err)
      os.Exit(1)
    }
    data := make([]byte, 4)
    conn.SetReadDeadline(time.Now().Add(time.Second * 10))
    _, err = conn.Read(data)
    if err != nil {
      fmt.Println("failed to read UDP msg because of ", err)
      continue
    }
    fmt.Println(data)
    num = num + 1
    if num > 9 {
      num = 0
    }
    time.Sleep(time.Second)
  }

  os.Exit(0)
}
