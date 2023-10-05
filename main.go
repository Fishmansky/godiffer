package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Fishmansky/godiffer/godiffer"
)

func passToCommand(cmd func(str1, str2 string) string) {
	args := os.Args[2:]
	if len(args) < 2 {
		fmt.Println("Not enough files passed!")
		os.Exit(1)
	}
	result := cmd(args[0], args[1])
	fmt.Println(result)
}

func main() {
	flagsCount := 0
	unionFlag := flag.Bool("u", false, "an union set")
	// complementFlag := flag.Bool("c", false, "a complement set")
	intersectionFlag := flag.Bool("i", false, "an intersection set")
	differenceFlag := flag.Bool("d", false, "a difference set")
	symetricdifferenceFlag := flag.Bool("s", false, "a symetric difference set")
	if *unionFlag {
		flagsCount++
	}
	// if *complementFlag {
	// 	flagsCount++
	// }
	if *intersectionFlag {
		flagsCount++
	}
	if *differenceFlag {
		flagsCount++
	}
	if *symetricdifferenceFlag {
		flagsCount++
	}
	// var outputFile string
	// flag.StringVar(&outputFile, "output", "", "an output file")
	flag.Parse()
	if flagsCount > 1 {
		fmt.Println("Cannot use more than one set type!")
		os.Exit(1)
	}
	if *unionFlag {
		passToCommand(godiffer.Union)
	}
	// if *complementFlag {
	// 	passToCommand(godiffer.Complement)
	// }
	if *intersectionFlag {
		passToCommand(godiffer.Intersection)
	}
	if *differenceFlag {
		passToCommand(godiffer.Difference)
	}
	if *symetricdifferenceFlag {
		passToCommand(godiffer.SymetricDifference)
	}
	passToCommand(godiffer.Difference)
}
