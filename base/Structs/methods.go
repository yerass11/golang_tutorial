package main

import "fmt"

type person struct {
	Id   int
	Name string
}

type account struct {
	Id   int
	Name string
	person
}

type MySlice []int

func (sl *MySlice) add(e int) {
	*sl = append(*sl, e)
}

func (sl MySlice) Count() int {
	return len(sl)
}

func (p *person) SetName(name string) {
	p.Name = name
}

func (acc *account) SetName(name string) {
	acc.Name = name
}

func (p person) UpdateName(name string) {
	p.Name = name
}

func main() {
	sl := MySlice{1, 2}
	sl.add(5)
	fmt.Println(sl.Count(), sl)

	fmt.Println("ais1ee")
	pers := person{1, "Azat"}
	fmt.Println(pers.Id, pers.Name)

	pers.SetName("Yerassyl")
	fmt.Println(pers.Id, pers.Name)

	pers.UpdateName("Zuck")
	fmt.Println(pers.Id, pers.Name)

	var acc account = account{
		Id:   1,
		Name: "ais1ee",
		person: person{
			Id:   2,
			Name: "Yerassyl Saiman",
		},
	}

	fmt.Println(acc.Id, acc.Name, acc.person.Id, acc.person.Name)
	acc.SetName("Mark Zuck")
	fmt.Println(acc.Id, acc.Name, acc.person.Id, acc.person.Name)
	acc.person.SetName("Elon Musk")
	fmt.Println(acc.Id, acc.Name, acc.person.Id, acc.person.Name)
}
