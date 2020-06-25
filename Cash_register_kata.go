/*
The new "Avengers" movie has just been released! There are a lot of
people at the cinema box office standing in a huge line. Each of them
has a single 100, 50 or 25 dollar bill. An "Avengers" ticket costs 25
dollars.

Vasya is currently working as a clerk. He wants to sell a ticket to
every single person in this line.

Can Vasya sell a ticket to every person and give change if he initially
has no money and sells the tickets strictly in the order people queue?

Return YES, if Vasya can sell a ticket to every person and give change
with the bills he has at hand at that moment. Otherwise return NO.

Examples:
Tickets([]int{25, 25, 50}) // => YES
Tickets([]int{25, 100}) // => NO. Vasya will not have enough money to
give change to 100 dollars
Tickets([]int{25, 25, 50, 50, 100}) // => NO. Vasya will not have the
right bills to give 75 dollars of change (you can't make two bills of
25 from one of 50)
*/

package main

import (
	"fmt"
	"sort"
)

const ticket_price int = 25

type CashRegister struct {
	cash  map[int]int // Notes we actually have
	notes []int       // Possible values of notes (must be sorted in descending order)
}

func NewCashRegister(notes []int) *CashRegister {
	var c CashRegister
	sort.Slice(notes, func(i, j int) bool { return notes[i] > notes[j] })
	c.notes = notes
	c.cash = make(map[int]int)
	for _, note := range c.notes {
		c.cash[note] = 0
	}
	return &c
}

func (c *CashRegister) Serve(price int, notes map[int]int) bool {
	have := 0
	for note, num := range notes {
		have += note * num
		c.cash[note] += num
	}
	if have < price {
		return false
	}
	change := have - price
	for change > 0 {
		for _, note := range c.notes {
			for c.cash[note] > 0 && note <= change {
				change -= note
				c.cash[note] -= 1
			}
		}
		if change > 0 {
			return false
		}
	}
	return true
}

func Tickets(peopleInLine []int) string {
	c := NewCashRegister([]int{100, 50, 25})
	for _, note := range peopleInLine {
		if !c.Serve(ticket_price, map[int]int{note: 1}) {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	c := NewCashRegister([]int{50, 100, 25})
	fmt.Println(c.Serve(25, map[int]int{25: 1}))
	fmt.Println(c.Serve(25, map[int]int{25: 1}))
	fmt.Println(c.Serve(25, map[int]int{50: 1}))
	fmt.Println(c.cash)
	fmt.Println(Tickets([]int{25, 25, 50}))
}
