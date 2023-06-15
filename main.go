package main

import "zuoye2/model"

func InitPerson(name string) model.Person {
	return model.Person{
		Name:  name,
		Level: 1,
		Exp:   0,
		Hp:    100,
		Atk:   10,
	}
}
func main() {
	person1 := InitPerson("亚索")
	person2 := InitPerson("石头人")
	person1.Attack(person2)
}
