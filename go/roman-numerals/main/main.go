package main

import (
	"bufio"
	"fmt"
	"os"
	"romannumerals"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	num, _ := strconv.Atoi(string(input))
	s, _ := romannumerals.ToRomanNumeral(num)
	fmt.Println(s)
}
