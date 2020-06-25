/*
Traveling Salesman kata

A traveling salesman has to visit clients. He got each client's address e.g.
"432 Main Long Road St. Louisville OH 43071" as a list.

The basic zipcode format usually consists of two capital letters followed by
a white space and five digits. The list of clients to visit was given as a
string of all addresses, each separated from the others by a comma, e.g. :

"123 Main Street St. Louisville OH 43071,432 Main Long Road St. Louisville
OH 43071,786 High Street Pollocksville NY 56432".

To ease his travel he wants to group the list by zipcode.

Task
The function travel will take two parameters r (addresses' list of all clients'
as a string) and zipcode and returns a string in the following format:

zipcode:street and town,street and town,.../house number,house number,...

The street numbers must be in the same order as the streets where they belong.

If a given zipcode doesn't exist in the list of clients' addresses return "zipcode:/"

Examples
r = "123 Main Street St. Louisville OH 43071,432 Main Long Road St. Louisville OH
43071,786 High Street Pollocksville NY 56432"

travel(r, "OH 43071") --> "OH 43071:Main Street St. Louisville,Main Long Road
St. Louisville/123,432"

travel(r, "NY 56432") --> "NY 56432:High Street Pollocksville/786"

travel(r, "NY 5643") --> "NY 5643:/"
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

const s string = "123 Main Street St. Louisville OH 43071, 432 Main Long Road St. Louisville OH 43071,786 High Street Pollocksville NY 56432, 54 Holy Grail Street Niagara Town ZP 32908"

func Travel(r, zipcode string) string {
	var answer string
	streets, houses := []string{}, []string{}
	answer = zipcode + ":"
	houseaddr, _ := regexp.Compile("[0-9]+")
	streetaddr, _ := regexp.Compile("[a-zA-Z].+")
	addresses := strings.Split(r, ",")
	for _, address := range addresses {
		if address[len(address)-8:] == zipcode {
			address = address[:len(address)-9]
			houses = append(houses, houseaddr.FindString(address))
			streets = append(streets, streetaddr.FindString(address))
		}
	}
	answer += strings.Join(streets, ",") + "/" + strings.Join(houses, ",")
	return answer
}

func main() {
	fmt.Println(Travel(s, "OH 43071"))
	fmt.Println(Travel(s, "OH 430"))

}
