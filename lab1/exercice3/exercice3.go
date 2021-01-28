package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	nextID := 101

	var (
		p   Person
		err error
	)
	nextID, err = inPerson(&p, nextID)
	if err != nil {
		fmt.Println("Invalid entry ... exiting")

	}
	printPerson(p)
	_ = nextID

}

//Person ...
type Person struct {
	lastName  string
	firstName string
	iD        int
}

func inPerson(p *Person, ID int) (int, error) {
	fmt.Print("Enter first name: ")
	reader := bufio.NewReader(os.Stdin)
	first, err1 := reader.ReadString('\n')
	fmt.Print("Enter last name: ")
	last, err2 := reader.ReadString('\n')
	*p = Person{first, last, ID}
	if err1 == nil || err2 == nil {
		return 0, err1
	}
	ID++
	return ID, nil

}

func printPerson(p Person) {
	f := p.firstName
	l := p.lastName
	i := p.iD
	fmt.Print("first name: ", f)
	fmt.Print("last name: ", l)
	fmt.Print("ID: ", i)

}
