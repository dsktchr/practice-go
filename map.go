package main

import "fmt"

func main() {
  var nilMap map[string]int
  fmt.Println(len(nilMap))
  fmt.Println(nilMap["A"])
  //panic!
  //nilMap["A"] = 10

  //マップの読み書き
  totalWinds := map[string]int{}
  totalWinds["国"] = 10
  totalWinds["給料"] = 100000
  totalWinds["インクリメント"] = 2
  totalWinds["インクリメント"]++
  fmt.Println(totalWinds)

  //集合を扱うことができる
  initSet := map[int]bool{}
  vals := []int{1, 20, 43, 1, 20, 3, 5, 100, 100, 33, 6}
  for _, v := range vals {
     initSet[v] = true
  }
  fmt.Println(len(vals), len(initSet))

  arr_1 := []int{1, 2, 3, 4, 2, 3}
  arr_2 := []int{2, 3, 6, 10, 10, 11}

  unionSet := union(makeSet(arr_1), makeSet(arr_2))
  fmt.Println(unionSet)
}

//集合を作成する
func makeSet[T] (arr []T) map[T]bool {
  newSet := map[T]bool{}
  for _, v := range arr {
     newSet[v] = true
  }
}

// 和集合を求める
func union[K] (set_1 map[K]bool, set_2 map[K]bool) map[K]bool {
  unionSet := make(map[K]bool)

  for k, v := range set_1 {
	  unionSet[k] = v
  }

  for k, v := range set_2 {
	  unionSet[k] = v
  }

  return unionSet
}
