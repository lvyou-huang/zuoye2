package model

import "fmt"

type Atk interface {
	Attack()
}
type Person struct {
	Name  string
	Level int
	Exp   int
	Hp    int
	Atk   int
}

func (person1 Person) Attack(person2 Person) {
	person2.Hp -= person1.Atk
	fmt.Printf("%v被%v攻击了,造成%v点伤害", person2.Name, person1.Name, person1.Atk)
}
