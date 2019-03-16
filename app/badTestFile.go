package main

import (
	"crypto/md5"
	"fmt"
)

func main() {

	var x int = 0
	y := 10
	if x > y {
		h := md5.New()
		fmt.Println(h)
	}

}
