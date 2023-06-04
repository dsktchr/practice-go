package main

import "fmt"

func main() {
    p := Person{
      "ダイスケ",
      "タチハラ",
      20,
    }
  
    fmt.Println(p.String())
}


type Person struct {
      LastName string
      FirstName string
      Age int
}

func (p Person) String() string {
  return fmt.Sprintf("%s %s : 年齢%d歳", p.LastName, p.FirstName, p.Age)
}
