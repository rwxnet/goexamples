package main
import (
  "fmt"
)

func main() {
  var a interface{}
  a = "abc"
  fmt.Println(a.(string) + "aaa")
  b := "abc"
  b = b + "def"

  fmt.Println(b)
}
