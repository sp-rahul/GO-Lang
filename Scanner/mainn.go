package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	names := make([]string, 20
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter name: ")
		scanner.Scan()

		txt := scanner.Text()
		if len(txt) != 0 {
			fmt.Println(txt)
			names = append(names, txt)

		} else {
			break
		}

	}
	fmt.Println(names)
}
