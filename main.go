package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Fishmansky/godiffer/godiffer"
)

func main() {
	flagsCount := 0
	unionFlag := flag.Bool("u", false, "an union set")
	intersectionFlag := flag.Bool("i", false, "an intersection set")
	differenceFlag := flag.Bool("d", false, "a difference set")
	symetricdifferenceFlag := flag.Bool("s", false, "a symetric difference set")
	if *unionFlag {
		flagsCount++
	}
	if *intersectionFlag {
		flagsCount++
	}
	if *differenceFlag {
		flagsCount++
	}
	if *symetricdifferenceFlag {
		flagsCount++
	}
	flag.Parse()
	if flagsCount > 1 {
		fmt.Println("Cannot use more than one set type!")
		os.Exit(1)
	}
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Not enough parameters passed!")
		os.Exit(1)
	}
	if *unionFlag {
		godiffer.Run(args[0], args[1], godiffer.UnionLines)
	}
	if *intersectionFlag {
		godiffer.Run(args[0], args[1], godiffer.IntersectionLines)
	}
	if *differenceFlag {
		godiffer.Run(args[0], args[1], godiffer.DifferenceLines)
	}
	if *symetricdifferenceFlag {
		godiffer.Run(args[0], args[1], godiffer.SymetricDifferenceLines)
	}
	godiffer.Run(args[0], args[1], godiffer.DifferenceLines)
}
