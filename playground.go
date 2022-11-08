package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var x string = `this is
my multiline
string`

	scanner := bufio.NewScanner(strings.NewReader(x))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
