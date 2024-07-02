package output

import "fmt"

func Redlog(s ...string) {
	fmt.Printf("\033[31m%s\033[0m\n", s)
}
