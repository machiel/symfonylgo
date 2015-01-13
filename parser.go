// package main
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func open(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	return file
}

func writeToFile(date string, line string, rootName string) {

	filename := rootName + "." + date + ".log"

	file := open(filename)
	defer file.Close()

	_, err := file.Write([]byte(line + "\n"))

	if err != nil {
		panic(err)
	}
}

func retrieveLines(filename string) []string {

	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	return lines
}

func printCriticalLines(lines []string) {

	r, rerr := regexp.Compile("\\.CRITICAL")
	e, eerr := regexp.Compile("\\.ERROR")

	if rerr != nil || eerr != nil {
		panic(rerr)
	}

	var critCount int
	var errCount int

	for _, line := range lines {
		if r.MatchString(line) {
			fmt.Println(line)
			critCount++
		}

		if e.MatchString(line) {
			fmt.Println(line)
			errCount++
		}
	}
	fmt.Printf("Error count: %d\n", errCount)
	fmt.Printf("Critical count: %d\n", critCount)
	fmt.Printf("Total lines: %d\n", len(lines))
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("usage: symfonylgo <command> <filename>")
		return
	}

	args := os.Args[1:]
	command := args[0]
	filename := args[1]

	fmt.Printf("Loading %s\n", filename)

	lines := retrieveLines(filename)

	fmt.Println(len(lines))

	if command == "critical" {
		printCriticalLines(lines)
		fmt.Println("Showing critical stuff")
		return
	}

	parts := strings.Split(filename, ".")

	var rootName string

	if len(parts) > 1 {
		rootName = parts[0]
	} else {
		rootName = filename
	}

	r, rerr := regexp.Compile("(([12][0-9]{3})-([01][0-9])-([0123][0-9]))")

	if rerr != nil {
		panic(rerr)
	}

	for _, line := range lines {
		occ := r.FindAllString(line, 2)
		if len(occ) > 0 {
			writeToFile(occ[0], line, rootName)
		}
	}

}
