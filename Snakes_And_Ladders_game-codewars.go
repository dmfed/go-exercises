/*
Codewars kata

Full description at
https://www.codewars.com/kata/587136ba2eefcb92a9000027/train/go

*/

package main

import (
	"fmt"
)

type SnakesLadders struct {
	swaps       map[int]int // These are "snakes" and "ladders"
	positions   map[int]int // 0 is player 1, 1 is player 2, the value of key is position of player
	turn        int         // Whose turn it is?
	running     bool        // Is the game still running?
	switchturns bool        // Should we switch turn? Turns to false if a player rolled double.
}

var sl *SnakesLadders

func (sl *SnakesLadders) NewGame() {
	sl.swaps = map[int]int{2: 38, 7: 14, 8: 31, 15: 26, 16: 6, 21: 42, 28: 84,
		36: 44, 46: 25, 49: 11, 51: 67, 62: 19, 64: 60,
		71: 91, 74: 53, 78: 98, 87: 94, 89: 68, 92: 88,
		95: 75, 99: 80}
	sl.positions = map[int]int{0: 0, 1: 0}
	sl.turn = 1 // This will switch to 0 at first call to Play()
	sl.running = true
	sl.switchturns = true
}

func (sl *SnakesLadders) Play(die1, die2 int) string {
	if !sl.running {
		return "Game over!"
	}
	if sl.switchturns { // Switching turn if we need to
		sl.turn = (sl.turn + 1) % 2
	}
	if die1 == die2 {
		sl.switchturns = false // Will not switch turns next time if player rolled double
	} else {
		sl.switchturns = true
	}
	dice := die1 + die2
	if sl.positions[sl.turn]+dice == 100 { // We have a winner
		sl.running = false
		return fmt.Sprintf("Player %v Wins!", sl.turn+1)
	}
	sl.positions[sl.turn] += dice
	if sl.positions[sl.turn] > 100 { // Overshot, needs to return
		sl.positions[sl.turn] -= (sl.positions[sl.turn] - 100) * 2
	}
	if swap, ok := sl.swaps[sl.positions[sl.turn]]; ok { // Check if we got to a "snake" or a "ladder"
		sl.positions[sl.turn] = swap
	}
	return fmt.Sprintf("Player %v is on square %v", sl.turn+1, sl.positions[sl.turn])
}
