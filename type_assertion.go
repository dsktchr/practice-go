package main

import (
  "fmt"
  "os"
)

type MyInt int

func main() {
  var i any
  var mine MyInt = 10
  i = mine
  i2 := i.(MyInt)
  fmt.Println(i2)
  
  // comma OK idiom
  i3, ok := i.(int)
  if !ok {
      err := fmt.Errorf("iの型（%v）が想定外です", i)
      fmt.Println(err.Error())
      os.Exit(1)
    }
  fmt.Println(i3)
}
