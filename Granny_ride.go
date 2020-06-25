/*
Your granny, who lives in town X0, has friends. These friends are
given in an array, for example: array of friends is

[ "A1", "A2", "A3", "A4", "A5" ].
The order of friends is this array must not be changed since this
order gives the order in which they will be visited.

These friends inhabit towns and you have an array with friends and
towns, for example:

[ ["A1", "X1"], ["A2", "X2"], ["A3", "X3"], ["A4", "X4"] ]
or
[ ("A1", "X1"), ("A2", "X2"), ("A3", "X3"), ("A4", "X4") ]
or
(C)
{"A1", "X1", "A2", "X2", "A3", "X3", "A4", "X4"}
which means A1 is in town X1, A2 in town X2... It can happen that we
don't know the town of one of the friends.

Your granny wants to visit her friends and to know how many miles she
will have to travel.

You will make the circuit that permits her to visit her friends. For
example here the circuit will contain:

X0, X1, X2, X3, X4, X0
and you must compute the total distance

X0X1 + X1X2 + .. + X4X0.
*/

package main

import (
	"fmt"
	"math"
)

func Tour(Friends []string, Towns map[string]string, distance map[string]float64) int {
	var (
		current_town   string
		previous_town  string
		total_distance float64
		radius         bool = true
	)
	for _, friend := range Friends {
		if val, ok := Towns[friend]; ok {
			previous_town = current_town
			current_town = val
		} else {
			continue
		}
		if radius {
			total_distance += distance[current_town]
			previous_town = current_town
			radius = false
		} else {
			total_distance += math.Sqrt(distance[current_town]*distance[current_town] - distance[previous_town]*distance[previous_town])
		}
		fmt.Println(total_distance)
	}
	total_distance += distance[current_town]
	return int(math.Floor(total_distance))
}
func main() {
	var f = []string{"A1", "A2", "A3", "A4", "A5"}
	var t = map[string]string{"A1": "X1", "A2": "X2", "A3": "X3", "A4": "X4"}
	var d = map[string]float64{"X1": 100.0, "X2": 200.0, "X3": 250.0, "X4": 300.0}
	fmt.Println(Tour(f, t, d))
}
