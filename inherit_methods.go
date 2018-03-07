package main

import "fmt"

type Base struct {
	id string
}

func (b *Base) Id() string {
	return b.id
}

func (b *Base) SetId(id string) {
	b.id = id
}

type Person2 struct {
	Base
	FirstName string
	LastName string
}

type Employee struct {
	Person2
	salary float64
}

func main()  {
	e := &Employee{Person2{Base{id:"123"}, "first", "last"}, 1290.123}
	fmt.Printf("e id: %s", e.Person2.Id())
	e.SetId("234")
	fmt.Printf("e id: %s", e.Person2.id)

}
//e id: 123e id: 234

