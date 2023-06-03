package main

import "fmt"

type person struct {
  name string //名前
  age int // 年齢
  pet string // ペット
}

func main() {
  julia := person{
	  name: "サンプル太朗",
	  age: 40,
	  pet: "cat",
       }

  fmt.Println(julia)
  fmt.Println(julia.name)
  fmt.Println(julia.age)
  fmt.Println(julia.pet)

  // 無名構造体
  pet := struct{
	name string
	kind string
  }{
	name: "プリン",
	kind: "dog",
  }
  fmt.Println(pet)
}
