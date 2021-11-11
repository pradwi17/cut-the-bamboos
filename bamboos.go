// 1st User will be able to input how many bamboos
// -> and then the user input how many joints for each of bamboos
// -> user will ask again to input how many times they want to cut the bamboos
// each cut will cut 1 of the bamboos joint
// user will see the process until the end of cycle cuts to see how many joints left for each bamboo

// Author : @Pradipta
// Dates  : 11/11/2021
// Purpose: bootcamp live coding session (golang)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("How many bamboos= ")
	scanner.Scan()
	var totalBamboos, _ = strconv.Atoi(scanner.Text())

	var bamboos = make([]int, totalBamboos) // assign user input with slice

	for i := 0; i < len(bamboos); i++ { // loop for user will able to inpur the joint for each bamboo
		fmt.Print("Bamboo Joint-", i+1, " = ")
		scanner.Scan()
		bamboos[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println("bamboos joint:", bamboos)

	fmt.Print("Cuts count: ") // input how many times cutting the bamboo
	scanner.Scan()
	var cuts, _ = strconv.Atoi(scanner.Text())

	fmt.Println("==========================")

	fmt.Println("Initial joint: ")
	for i := 0; i < len(bamboos); i++ { // loop for every bamboo
		fmt.Print("|")
		for j := 0; j < bamboos[i]; j++ { // loop on how many bamboo's joint
			fmt.Print("-") // print the joint for simulate
		}
		fmt.Println("")
	}

	for i := 0; i < cuts; i++ { // loop for every cuts
		fmt.Println("==========================")
		fmt.Println("Cycle Cut:", i+1)
		for j := 0; j < len(bamboos); j++ { // loop on every bamboo
			fmt.Print("|")

			bamboos[j] -= 1 // cut bamboo (-1) for every cut (bamboo[j] = bamboo[j] - 1)

			for k := 0; k < bamboos[j]; k++ { // loop joint for very bamboo
				fmt.Print("-") // simulate
			}

			fmt.Println("")
		}
	}

}
