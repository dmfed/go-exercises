/*
This Kata is like the game of Snakes & Ladders

There is an array representing the squares on the game board.

The starting square is at array element 0. The final square is the last array element.

At each "turn" you move forward a number of places (according to the next dice throw).

The value at the square you end up on determines what happens next:

0 Stay where you are (until next turn)
+n This is a "ladder". Go forward n places
-n This is a "snake". Go back n places
Each snake or ladder will always end on a square with a 0, so you will only go up or 
down one at a time.

Rules
You are given a number of dice throws. The game continues until either:
You have no throws left, OR
You end up exactly on the final square
At each turn, make your move, then go up the "ladders" and down the "snakes" as appropriate.
If the dice roll overshoots the final square then you cannot move. Roll the dice again.
Task
Return the index of the array element that you ended up on at the end of the game.
*/

package main

func SnakesAndLadders(board, dice []int) int {
	end := len(board) - 1
	pos := 0
	for roll := range dice {
		if newpos := pos + roll + board[pos+roll]; newpos < end {
			pos = newpos
		} else if newpos = end {
			return newpos
		} 
	}
  	return pos
}