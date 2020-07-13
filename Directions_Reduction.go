/*
 * Once upon a time, on a way through the old wild mountainous west,…
… a man was given directions to go from one point to another. The
directions were "NORTH", "SOUTH", "WEST", "EAST". Clearly "NORTH" and
"SOUTH" are opposite, "WEST" and "EAST" too.

Going to one direction and coming back the opposite direction right away
is a needless effort. Since this is the wild west, with dreadfull
weather and not much water, it's important to save yourself some energy,
otherwise you might die of thirst!

How I crossed a mountain desert the smart way.
The directions given to the man are, for example, the following
(depending on the language):

["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"].

In ["NORTH", "SOUTH", "EAST", "WEST"], the direction "NORTH" + "SOUTH"
is going north and coming back right away.

The path becomes ["EAST", "WEST"], now "EAST" and "WEST" annihilate each
other, therefore, the final result is [] (nil in Clojure).

In ["NORTH", "EAST", "WEST", "SOUTH", "WEST", "WEST"], "NORTH" and
"SOUTH" are not directly opposite but they become directly opposite
after the reduction of "EAST" and "WEST" so the whole path is reducible
to ["WEST", "WEST"].

Task
Write a function dirReduc which will take an array of strings and
returns an array of strings with the needless directions removed
(W<->E or S<->N side by side).
*/

package main

import (
	"fmt"
)

func DirReduc(directions []string) []string {
	var result []string
	for _, direction := range directions {
		if len(result) > 0 && is_opposite(direction, result[len(result)-1]) {
			result = result[:len(result)-1]
		} else {
			result = append(result, direction)
		}
	}
	return result
}

func is_opposite(a, b string) bool {
	result := false
	if (a == "NORTH" && b == "SOUTH") || (a == "SOUTH" && b == "NORTH") {
		result = true
	} else if (a == "EAST" && b == "WEST") || (a == "WEST" && b == "EAST") {
		result = true
	}
	return result
}

func main() {
	fmt.Println(DirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}))
}
