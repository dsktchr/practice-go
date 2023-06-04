package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("従業員: %s, ID=%s", e.Name, e.ID)
}

type Manager struct {
	//埋め込み型
	// 名前を付けない場合は、埋め込み型となり、Employee構造体が
	// Manager構造体に昇格する
	Employee // 型のみ
	Reposrt []Employee
}

func (m Manager) FindEmployees() []Employee {
	newEmployees := []Employee {
		Employee {
			"ベネット",
			"100",
		},
		Employee {
			"行秋",
			"200",
		},
	}

	return newEmployees
}

func main() {
	// マネージャー
	m := Manager {
		Employee: Employee {
			"クレー",
			"300",
		},
		Reposrt: []Employee{},
	}
	
	//クレーのIDが表示される
	fmt.Println(m.ID)

	// クレーの詳細情報が表示される
	fmt.Println(m.Description())

	// ガチャで2引きする
	m.Reposrt = m.FindEmployees()

	fmt.Println(m.Employee)
	fmt.Println(m.Reposrt)
}
