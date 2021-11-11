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

	var bamboos = make([]int, totalBamboos)

	for i := 0; i < len(bamboos); i++ {
		fmt.Print("Bamboo Joint-", i+1, " = ")
		scanner.Scan()
		bamboos[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println("bamboos joint:", bamboos)

	fmt.Print("Cuts count: ")
	scanner.Scan()
	var cuts, _ = strconv.Atoi(scanner.Text())

	fmt.Println("==========================")

	fmt.Println("Initial joint: ")
	for i := 0; i < len(bamboos); i++ {
		fmt.Print("|")
		for j := 0; j < bamboos[i]; j++ {
			fmt.Print("-")
		}
		fmt.Println("")
	}

	for i := 0; i < cuts; i++ {
		fmt.Println("==========================")
		fmt.Println("Cycle Cut:", i+1)
		for j := 0; j < len(bamboos); j++ {
			fmt.Print("|")

			bamboos[j] -= 1

			for k := 0; k < bamboos[j]; k++ {
				fmt.Print("-")
			}

			fmt.Println("")
		}
	}

}
