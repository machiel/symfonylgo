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

func main() {

	filename := "prod.log"
	parts := strings.Split(filename, ".")

	var rootName string

	if len(parts) > 1 {
		rootName = parts[0]
	} else {
		rootName = filename
	}

	fmt.Printf("Loading %s\n", filename)

	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	fmt.Println(len(lines))

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
