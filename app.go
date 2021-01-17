package main

import (
	"bufio"
	"fmt"
	"os"

	m "github.com/ajengpu/roman-numbers/model"
)

func main() {
	var args []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			args = append(args, text)
		} else {
			break
		}
	}

	if scanner.Err() != nil {
		fmt.Printf("%v", scanner.Err())
	}

	for _, a := range args {
		fmt.Printf(m.ExecuteCommad(a))
	}
}
