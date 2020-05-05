// https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3/train/go

package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	var names = strings.Split(strings.ToUpper(name), " ")
	return fmt.Sprintf("%c.%c", names[0][0], names[1][0])
}

func main() {
	fmt.Println(AbbrevName("Sam Harris"))
	fmt.Println(AbbrevName("Patrick Feeney"))
}
