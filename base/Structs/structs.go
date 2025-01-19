package Structs

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	string
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	fmt.Println("ais1ee")

	var acc Account = Account{
		Id: 1,
		//Name:  "ais13",
		Owner: Person{Name: "Yerass"},
	}
	fmt.Printf("%#v\n", acc)

	acc.Owner = Person{2, "Yerassyl Saiman", "Islam Karima 70k1/10c, Almaty"}
	fmt.Printf("%#v\n", acc)
	fmt.Println(acc.Owner.Address)

	fmt.Println(acc.Name)
	fmt.Println(acc.Owner.Name)
}
