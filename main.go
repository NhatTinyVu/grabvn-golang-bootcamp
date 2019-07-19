package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("> ")
	
		for scanner.Scan(){
		text := scanner.Text()

		if result, err := eval(text); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}

		fmt.Println("> ")
	}

}