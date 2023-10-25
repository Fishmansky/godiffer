package godiffer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Contains(arr []byte, b byte) bool {
	for _, v := range arr {
		if v == b {
			return true
		}
	}
	return false
}

func ContainsLine(arr []string, s string) bool {
	for _, l := range arr {
		if l == s {
			return true
		}
	}
	return false
}

func Run(f1, f2 string, fn func(arr1, arr2 []string) string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	file1, err := os.Open(path + "/" + f1)
	if err != nil {
		log.Println(err)
	}
	file2, err := os.Open(path + "/" + f2)
	if err != nil {
		log.Println(err)
	}
	defer file1.Close()
	defer file2.Close()
	f1arr := []string{}
	f2arr := []string{}
	s1 := bufio.NewScanner(file1)
	for s1.Scan() {
		f1arr = append(f1arr, s1.Text())
	}
	s2 := bufio.NewScanner(file2)
	for s2.Scan() {
		f2arr = append(f2arr, s2.Text())
	}
	result := fn(f1arr, f2arr)
	fmt.Println(result)
	os.Exit(0)
}

func Union(str1, str2 string) string {
	strArr1 := []byte(str1)
	strArr2 := []byte(str2)
	result := []byte{}
	for _, v := range strArr1 {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	for _, v := range strArr2 {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return string(result)
}

func UnionLines(of1, of2 []string) string {
	var result = []string{}
	for _, l := range of1 {
		if !ContainsLine(result, l) {
			result = append(result, l)
		}
	}
	for _, l := range of2 {
		if !ContainsLine(result, l) {
			result = append(result, l)
		}
	}
	return strings.Join(result, "\n")
}

func Intersection(str1, str2 string) string {
	strArr1 := []byte(str1)
	strArr2 := []byte(str2)
	result := []byte{}
	for _, v := range strArr1 {
		if Contains(strArr2, v) {
			result = append(result, v)
		}
	}
	return string(result)
}

func IntersectionLines(of1, of2 []string) string {
	var result = []string{}
	for _, l := range of1 {
		if ContainsLine(of2, l) {
			result = append(result, l)
		}
	}
	return strings.Join(result, "\n")
}

func Difference(str1, str2 string) string {
	strArr1 := []byte(str1)
	strArr2 := []byte(str2)
	result := []byte{}
	for _, v := range strArr1 {
		if !Contains(strArr2, v) {
			result = append(result, v)
		}
	}
	return string(result)
}

func DifferenceLines(of1, of2 []string) string {
	result := []string{}
	for _, l := range of1 {
		if !ContainsLine(of2, l) {
			result = append(result, l)
		}
	}
	return strings.Join(result, "\n")
}

func SymetricDifference(str1, str2 string) string {
	strArr1 := []byte(str1)
	strArr2 := []byte(str2)
	result := []byte{}
	for _, v := range strArr1 {
		if !Contains(strArr2, v) {
			result = append(result, v)
		}
	}
	for _, v := range strArr2 {
		if !Contains(strArr1, v) {
			result = append(result, v)
		}
	}
	return string(result)
}

func SymetricDifferenceLines(of1, of2 []string) string {
	result := []string{}
	for _, l := range of1 {
		if !ContainsLine(of2, l) {
			result = append(result, l)
		}
	}
	for _, l := range of2 {
		if !ContainsLine(of1, l) {
			result = append(result, l)
		}
	}
	return strings.Join(result, "\n")
}
