package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Hello go")
	fmt.Println(nextLine())
	fmt.Printf("%s -> %d", "10", s2i("10"))
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}
