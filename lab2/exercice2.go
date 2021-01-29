package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	show1 := Play{"Tartuffe", make([]*Ticket, 25), time.Date(2020, time.March, 3, 22, 0, 0, 0, time.UTC)}
	show2 := Play{"Macbeth", make([]*Ticket, 25), time.Date(2020, time.April, 10, 19, 0, 0, 0, time.UTC)}
	theater := NewTheater(32, 5, []Play{show1, show2})
	fmt.Println("--------------------------------------------")
	fmt.Println("press 'b' to buy a ticket or 'q' to quit: ")
	fmt.Println("--------------------------------------------")
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')

	for {

		if text == "q\n" {
			break
		}
		var ticket Ticket
		fmt.Println("Enter Your Name: ")
		fmt.Print("-> ")
		name, _ := reader.ReadString('\n')
		ticket.customer = name
		fmt.Println("Pick a show: Tartuffe or Macbeth")
		fmt.Print("-> ")
		show, _ := reader.ReadString('\n')
		if strings.EqualFold(show, theater.shows[0].GetName()+"\n") {
			ticket.show = &show1
		} else if strings.EqualFold(show, theater.shows[1].GetName()+"\n") {
			ticket.show = &show2
		} else {
			break
		}
		fmt.Println("Enter the seat's number: ")
		fmt.Print("-> ")
		number, _ := reader.ReadString('\n')
		seat, _ := strconv.Atoi(number)
		s := Seat{seat}
		ticket.siege = &s
		ticket.show.purchased = append(ticket.show.purchased)

		fmt.Println("--------------------------------------------")
		fmt.Println("press 'b' to buy a ticket or 'q' to quit: ")
		fmt.Println("--------------------------------------------")
		fmt.Print("-> ")
		text, _ = reader.ReadString('\n')

	}
}

//Seat is...
type Seat struct {
	number int
}

//Ticket is...
type Ticket struct {
	customer string
	siege    *Seat
	show     *Play
}

//Show is...
type Show interface {
	GetName() string
	GetShowStart() time.Time
	AddPurchase(*Ticket) bool
	IsNotPurchased(*Ticket) bool
}

//Play is...
type Play struct {
	name      string
	purchased []*Ticket
	date      time.Time
}

// GetName is ..
func (p Play) GetName() string {
	return p.name
}

//GetShowStart is ..
func (p Play) GetShowStart() time.Time {
	return p.date
}

//AddPurchase is ..
func (p Play) AddPurchase(t *Ticket) bool {
	if !p.IsNotPurchased(t) {
		return false
	}
	p.purchased = append(p.purchased, t)
	return true
}

//IsNotPurchased is..
func (p Play) IsNotPurchased(t *Ticket) bool {
	if len(p.purchased) == 0 {
		return true
	}
	for _, i := range p.purchased {
		if i.siege.number == t.siege.number && i.show == t.show {
			return false
		}
	}
	return true

}

//Theater is...
type Theater struct {
	seats []Seat
	shows []Play
	rows  int32
}

//NewSeat is...
func NewSeat(number int) Seat {
	return Seat{number}
}

//NewTicket is...
func NewTicket(name string, seat *Seat, show *Play) Ticket {
	return Ticket{name, seat, show}
}

//NewTheater is...
func NewTheater(numberseats int32, numberrows int32, shows []Play) Theater {
	return Theater{make([]Seat, numberseats), shows, numberrows}
}

func (p *Play) emptySeat(ticket *Ticket, temp int) Ticket {
	if p.IsNotPurchased(ticket) {
		fmt.Println("Ticket bought!")
		return *ticket
	}
	for _, i := range p.purchased {
		if i.siege.number != temp {
			ticket.siege.number = temp
			fmt.Printf("The seat chosen is taken. We changed it to number: %v", temp)
			return *ticket
		}
	}
	temp++
	p.emptySeat(ticket, temp)
	return *ticket

}
