package main

import "fmt"

func main() {
  x := []int{1, 5: 4, 6, 10: 100, 15}
  fmt.Println(x)

  fmt.Println(x[0])

  x[0] = 777

  fmt.Println(x)

  var slice_1 []int

  fmt.Println(slice_1 == nil)

  slice_1 = append(slice_1, 100)

  slice_2 := []int{20, 30, 40}
  slice_1 = append(slice_1, slice_2...)
  fmt.Println(slice_1)
  
  make_slice_1 := make([]string, 3)
  fmt.Println(make_slice_1, len(make_slice_1), cap(make_slice_1))

  make_slice_1 = append(make_slice_1, "hello", "world", "japanese")
  //最後に追加されてしまう
  fmt.Println(make_slice_1)

  //長さ0, キャパシティ4のスライスを作成する
  make_slice_2 := make([]int, 0, 4)
  fmt.Println(make_slice_2, len(make_slice_2), cap(make_slice_2))

  make_slice_2 = append(make_slice_2, 2, 4, 6, 8)
  fmt.Println(make_slice_2)
}
