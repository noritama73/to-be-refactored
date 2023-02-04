package a

import "fmt"

func f(j int) {             // want "args: 1"
	for i := 0; i < 10; i++ { // want "found for"
		if i % 2 == 0 {         // want "found if"
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}                           // want "lines: 9"
