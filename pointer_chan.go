package main
import (
    "fmt"
    "sync"
    "net"
    "time"
)
type ReplyMsg struct {
  Buffer []byte
  SippAddr *net.UDPAddr
}

func HandleRecv(c chan *ReplyMsg, wg *sync.WaitGroup) {

  var rm  *ReplyMsg
  for {
    select {
	    case rm = <-c:
		    fmt.Println("Receive Message")
		    fmt.Println(rm.Buffer)
		    fmt.Println(rm.SippAddr)
    }
  }
}

func main() {
    c := make(chan *ReplyMsg, 1) // at least 1, otherwise main goroutine will block
    wg := sync.WaitGroup{}
    wg.Add(1)
    aa := make([]*ReplyMsg, 1, 2)
    aa[0] = new(ReplyMsg)
    aa[0].Buffer = []byte("abc")
    aa[0].SippAddr, _ = net.ResolveUDPAddr("udp", "127.0.0.1" + ":" + "0")
    c <- aa[0]
    time.Sleep(10 * time.Second)
    aa = aa[:0]
    fmt.Println(aa)
    go HandleRecv(c, &wg)
    wg.Wait()
}
