package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	s, _ := strconv.Unquote(string(bytes))
	fmt.Println(string(s))
}
