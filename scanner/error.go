package scanner

import "fmt"

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %v] Error %v : %v\n", line, where, message)
}
