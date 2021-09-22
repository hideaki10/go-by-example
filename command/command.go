package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]

	// arg := os.Args[3]

	// fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)
	// fmt.Println(arg)

	// wordPtr := flag.String("xargs", "foo", "a string")

	// numPtr := flag.Int("numb", 43, "an int")
	// boolPtr := flag.Bool("fork", false, "a bool")

	// var svar string
	// flag.StringVar(&svar, "svar", "bar", "a string var")

	// flag.Parse()

	// fmt.Println("word string: ", *wordPtr)
	// fmt.Println("word num : ", *numPtr)
	// fmt.Println("word bool : ", *boolPtr)
	// fmt.Println("word svar : ", svar)
	// fmt.Println("word: flag ", flag.Args())

	fooCmd := flag.NewFlagSet("xargs", flag.ExitOnError)
	//fooName := fooCmd.String("l", "dd", "the position of args")
	fooName := fooCmd.String("l", "level", "level")

	barCmd := flag.NewFlagSet("bars", flag.ExitOnError)
	barLevel := barCmd.String("l", "level", "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands ")
		os.Exit(1)
	}
	switch os.Args[1] {

	case "xargs":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'xargs'")
		fmt.Println("  level:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bars":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bars'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
