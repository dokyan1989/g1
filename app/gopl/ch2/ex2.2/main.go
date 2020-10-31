package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/dokyan1989/g1/app/gopl/ch2/tempconv"
)

func main() {
	var inputs []string
	if len(os.Args[1:]) > 0 {
		inputs = os.Args[1:]
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			inputs = append(inputs, input.Text())
		}
	}

	for _, in := range inputs {
		t, err := strconv.ParseFloat(in, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

}
