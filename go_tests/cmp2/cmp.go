package cmp 

import "time"

type Person struct {
	Name string
	Age int
	DateAdded time.Time
}

// ユーザー作成
func CreatePerson(name string, age int) Person {
	return Person{
		Name: name,
		Age: age,
		DateAdded: time.Now(),	
	}
}
