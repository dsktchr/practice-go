package main

import "fmt"


func main() {
  var x int = 10
  var y string = "HelloWorld"

  pointerX := &x
  pointerY := &y

  // ポインタを直接参照する
  fmt.Println(pointerX)
  fmt.Println(pointerY)
  // ポインタを間接参照する
  fmt.Println(*pointerX)
  fmt.Println(*pointerY)

  // 間接参照した値と新しい値を組み合わせる
  z := *pointerY + ", piyopiyo"
  fmt.Println(z)
}
